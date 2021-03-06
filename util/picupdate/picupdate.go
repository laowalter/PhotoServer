package main

/* Generate thumbnail and store to mongodb
Database: album
Collection: pic
*/
import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"flag"
	"fmt"
	"image/jpeg"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/barasher/go-exiftool"
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
	file, err := os.OpenFile("picupdate.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.Print("Logging to a file in Go!")

	pathPtr := flag.String("path", ".", "Input a new path")
	flag.Parse()

	absPath, err := filepath.Abs(*pathPtr)
	if err != nil {
		fmt.Printf("Can not get the Absolut path of: %v", Red(*pathPtr))
		return
	}

	files, err := global.ListFiles(absPath, global.PIC) //仅处理图像找到截图视频exif信息再说。
	if err != nil {
		panic(err)
	}

	if len(files) == 0 {
		fmt.Printf("No files in directory: %v need to update!\n", Red(absPath))
		return
	}

	err = insert(files)
	if err != nil {
		fmt.Println("Insert Error")
		panic(err)
	}

}

func insert(files []string) error {

	total := len(files)
	fmt.Printf("Total files number is: %v\n", total)
	database, collection, uri := "album", "pic", "mongodb://localhost:27017"

	db, _ := connectToDB(uri, database)
	col := db.Collection(collection)

	{ //Create index by md5
		mod := mongo.IndexModel{
			Keys: bson.M{
				"md5": 1,
			}, Options: options.Index().SetUnique(true),
		}
		_, err := col.Indexes().CreateOne(context.TODO(), mod)
		if err != nil { // Check if the CreateOne() method returned any errors
			fmt.Println("Indexes().CreateOne() ERROR:", err)
			os.Exit(1)
		} else { // API call returns string of the index name
			fmt.Printf("Notice: Use %s of file as the unique index.\n", Red("md5sum"))
		}

		mod = mongo.IndexModel{
			Keys: bson.M{
				"path": 1,
			}, Options: options.Index().SetUnique(true),
		}
		_, err = col.Indexes().CreateOne(context.TODO(), mod)
		if err != nil { // Check if the CreateOne() method returned any errors
			fmt.Println("Indexes().CreateOne() ERROR:", err)
			os.Exit(1)
		} else { // API call returns string of the index name
			fmt.Printf("Notice: Use %s of file as the unique index.\n", Red("md5sum"))
		}

	}

	var _elapseTime time.Duration
	start := time.Now()

	for index, file := range files {
		fmt.Printf("Dealing with %v/%v......", index+1, total)
		document := &global.Document{
			FileName:    filepath.Base(file),
			Path:        file,
			ContentType: "jpeg",
			ImportTime:  time.Now(),
		}

		var err error
		document.Thumbnail, err = thumb(file)
		if err != nil {
			continue
		}

		document.Md5, err = fileMd5(document.Path)
		if err != nil {
			continue
		}

		document.Exif, document.GPSPosition = exif(file)

		if document.CreateDate.IsZero() {
			log.Printf("[%v]: CreateDate of %v is empty\n", time.Now(), file)
		}
		if len(document.FullImageSize) == 0 {

			log.Printf("[%v]: Image Size of %v is empty\n", time.Now(), file)
		}

		_, insertErr := col.InsertOne(context.TODO(), document)
		if insertErr != nil {
			//check if the Md5 already exist in db is yes
			filter := bson.M{"md5": document.Md5}
			count, err := col.CountDocuments(context.TODO(), filter)
			if err != nil {
				fmt.Println("Can not search by Md5")
				return err
			}
			if count >= 1 {
				update := bson.M{"$set": bson.M{"filename": document.FileName, "path": document.Path}}
				_ = col.FindOneAndUpdate(context.TODO(), filter, update)
				fmt.Printf("File: %v already exist, filename and path updated.\n", Red(document.Path))
			} else {
				fmt.Println(Red("Md5 is not unique index in database! Wrong"))
				return err
			}

		} else {
			fmt.Printf("%s %v", Green("Inserted"), document.FileName)
		}
		_elapseTime = time.Since(start)
		elapseTime := int64(_elapseTime) / int64(index+1) * int64(total-index-1)
		fmt.Printf("...%v Seconds left.\n", elapseTime/int64(time.Second))
	}
	return nil
}

func exif(file string) (global.Exif, global.GPSPosition) {

	var exifInfo global.Exif
	var gpsInfo global.GPSPosition

	et, err := exiftool.NewExiftool()
	if err != nil {
		fmt.Printf("Error when intializing: %v\n", err)
		return exifInfo, gpsInfo
	}
	defer et.Close()

	fileInfos := et.ExtractMetadata(file)
	for _, fileInfo := range fileInfos {
		if fileInfo.Err != nil {
			fmt.Printf("Error concerning %v: %v\n", fileInfo.File, fileInfo.Err)
			continue
		}

		for k, v := range fileInfo.Fields {
			//fmt.Printf("KEY:%v-----VALUE:%v\n", k, v)
			switch k {
			case "CreateDate":
				_cDate := fmt.Sprintf("%v", v)
				exifInfo.CreateDate = createDate(_cDate)
				if exifInfo.CreateDate.IsZero() {
					log.Printf("[%v]: CreateDate of %v is empty\n", time.Now(), file)

				}

			case "Make":
				exifInfo.Make = fmt.Sprintf("%v", v)
			case "Model":
				exifInfo.Model = fmt.Sprintf("%v", v)
			case "LensSpec":
				exifInfo.LensSpec = fmt.Sprintf("%v", v)
			case "LensID":
				exifInfo.LensID = fmt.Sprintf("%v", v)
			case "ShutterSpeed":
				exifInfo.ShutterSpeed = fmt.Sprintf("%v", v)
			case "ExposureTime":
				exifInfo.ExposureTime = fmt.Sprintf("%v", v)
			case "ISO":
				exifInfo.ISO = fmt.Sprintf("%v", v)
			case "Aperture":
				exifInfo.Aperture = "f/" + fmt.Sprintf("%v", v)
			case "ExposureCompensation":
				exifInfo.ExposureCompensation = fmt.Sprintf("%v", v)
			case "FocalLength":
				exifInfo.FocalLength = fmt.Sprintf("%v", v)
			case "GPSPosition":
				_gpsPosition := fmt.Sprintf("%v", v)
				gps := strings.Split(_gpsPosition, ",")
				gpsInfo.Latitude, gpsInfo.Longitude = gps[0], gps[1]
			case "FullImageSize":
				exifInfo.FullImageSize = fmt.Sprintf("%v", v)
			case "ImageSize":
				exifInfo.FullImageSize = fmt.Sprintf("%v", v)
			}
		}

	}
	return exifInfo, gpsInfo
}

func thumb(file string) (string, error) { //file: wholepath contains filename.
	runtime.GOMAXPROCS(runtime.NumCPU())
	img, err := imaging.Open(file, imaging.AutoOrientation(true))
	if err != nil {
		fmt.Printf("File: %s Can not open\n", file)
		return "", err
	}

	rectangle := img.Bounds()
	t_width := int(rectangle.Dx() * global.ThumbnailHeight / rectangle.Dy())
	thumb := imaging.Thumbnail(img, t_width, global.ThumbnailHeight, imaging.CatmullRom)

	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, thumb, nil)
	if err != nil {
		fmt.Printf("jpeg Encode error on %s\n", file)
		return "", err
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
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

func createDate(ct string) time.Time {
	_time, err := time.Parse("2006:01:02 15:04:05", ct)
	if err != nil {
		_time, err = time.Parse("2006:01:02 15:04:05-07:00", ct)
		if err != nil {
			_ct := []rune(ct)
			if string(_ct[len(_ct)-2:]) == "下午" {
				ct = string(_ct[:len(_ct)-2]) + "PM"
				fmt.Println(ct)
			} else {
				ct = string(_ct[:len(_ct)-2]) + "AM"
				fmt.Println(ct)
			}
			_time, err = time.Parse("2006:01:02 15:04:05PM", ct)
			if err != nil {
				return time.Time{}

			}
		}
	}

	return _time
}
