package main

/* Generate thumbnail and store to mongodb
Database: album
Collection: pic
*/
import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"image/jpeg"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/disintegration/imaging"
	. "github.com/logrusorgru/aurora"
	"github.com/photoServer"
	"github.com/photoServer/myimage"
	"github.com/photoServer/mymongo"
)

func main() {
	database, collection := "album", "pic"
	client, ctx := mymongo.ConnectDB("mongodb://localhost:27017")
	db := &mymongo.DBconn{Client: client, Ctx: ctx, DBName: database}

	path := "/home/walter/go/src/github.com/photoServer/testData"

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
			panic(err)
		}

		err = db.InsertOne(collection, document)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s %v\n", Green("Inserted"), document.FileName)

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
