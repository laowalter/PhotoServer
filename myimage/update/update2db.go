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
	"flag"
	"fmt"
	"image/jpeg"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/disintegration/imaging"
	"github.com/photoServer/global"

	. "github.com/logrusorgru/aurora"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectToDB(uri, dbname string) (*mongo.Database, error) {

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("mongo.Connect() ERROR:", err)
		os.Exit(1)
	}
	return client.Database(dbname), nil
}

func main() {
	pathPtr := flag.String("path", ".", "Input a new path")
	flag.Parse()

	files := global.ListPics(*pathPtr) //仅处理图像找到截图视频exif信息再说。
	total := len(files)

	if total == 0 {
		fmt.Printf("No files in directory: %v need to update!\n", Red(*pathPtr))
		return
	}

	fmt.Printf("Total files number is: %v\n", total)
	database, collection, uri := "album", "pic", "mongodb://localhost:27017"

	db, _ := connectToDB(uri, database)
	col := db.Collection(collection)

	{ //Create index by md5
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
	}

	i := 0
	for _, file := range files {

		i++
		fmt.Printf("Dealing with %v/%v......", i, total)
		document := &global.Document{
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
			if count >= 1 {
				update := bson.M{"$set": bson.M{"filename": document.FileName, "path": document.Path}}
				_ = col.FindOneAndUpdate(context.TODO(), filter, update)
				fmt.Printf("File: %v already exist, filename and path updated.\n", Red(document.Path))
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
