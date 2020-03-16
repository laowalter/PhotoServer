package main

/* Generate thumbnail and store to mongodb
Database: album
Collection: pic
*/
import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"image/jpeg"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/disintegration/imaging"
	"github.com/photoServer"
	"github.com/photoServer/myimage"

	. "github.com/logrusorgru/aurora"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	database, collection, uri := "album", "pic", "mongodb://localhost:27017"
	path := "/home/walter/go/src/github.com/photoServer/testData"

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("mongo.Connect() ERROR:", err)
		os.Exit(1)
	}
	col := client.Database(database).Collection(collection)

	mod := mongo.IndexModel{
		Keys: bson.M{
			"md5": 1, // index in ascending order, md5 is from the `bson: "md5"`
		}, Options: options.Index().SetUnique(true),
	}
	ind, err := col.Indexes().CreateOne(context.TODO(), mod)
	if err != nil { // Check if the CreateOne() method returned any errors
		fmt.Println("Indexes().CreateOne() ERROR:", err)
		os.Exit(1)
	} else { // API call returns string of the index name
		fmt.Println("CreateOne() index:", ind)
	}
	files := myimage.ListPics(path)
	total := len(files)
	fmt.Printf("Total files number is: %v\n", total)

	i := 0
	for _, file := range files {

		i++
		fmt.Printf("Dealing with %v/%v......", i, total)
		document := &photoServer.Document{
			FileName:    filepath.Base(file),
			Path:        file,
			CreateTime:  time.Now(),
			ContentType: "jpeg",
			Thumbnail:   thumb(file),
		}

		var err error
		document.Md5, err = fileMd5(document.Path)
		if err != nil {
			fmt.Println("Md5 Error", document.Path)
			os.Exit(1)
		}

		_, insertErr := col.InsertOne(context.TODO(), document)
		if insertErr != nil {
			//check if the Md5 already exist in db is yes
			filter := bson.M{"md5": document.Md5}
			count, err := col.CountDocuments(context.TODO(), filter)
			if err != nil {
				log.Fatal(err)
			}
			if count == 1 {
				fmt.Printf("File: %v already exist\n", Red(document.Path))

			} else {
				fmt.Println(Red("Md5 is not unique index in database! Wrong"))
				os.Exit(1)
			}

		} else {
			fmt.Printf("%s %v\n", Green("Inserted"), document.FileName)

		}

	}
}

func thumb(file string) []byte { //file: wholepath contains filename.
	runtime.GOMAXPROCS(runtime.NumCPU())
	img, err := imaging.Open(file)
	if err != nil {
		panic(err)
	}

	thumb := imaging.Thumbnail(img, 200, 200, imaging.CatmullRom)
	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, thumb, nil)
	if err != nil {
		panic(err)
	}

	return buf.Bytes()
}

func fileMd5(filePath string) (string, error) {
	var returnMD5String string
	file, err := os.Open(filePath)
	if err != nil {
		return returnMD5String, err
	}
	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String, err
	}
	hashInBytes := hash.Sum(nil)[:16]
	returnMD5String = hex.EncodeToString(hashInBytes)
	return returnMD5String, nil

}