package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	. "github.com/logrusorgru/aurora"
	"github.com/photoServer/global"
	"github.com/photoServer/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func main() {

	documentList, err := model.QueryAllPhotosOnepage()
	if err != nil {
		panic(err)
	}

	col, err := connectToDB(global.MongoUri, global.DBname, "gps")
	if err != nil {
		panic(err)
	}

	fmt.Println("Total number of photos is:", len(documentList), "Scanning...")

	counter := 0
	for _, document := range documentList {

		strings.TrimSpace(document.Latitude)
		strings.TrimSpace(document.Longitude)

		if document.Latitude == "" || document.Longitude == "" {
			continue
		}

		fmt.Printf("Dealing with: %v, latitude: %v, longitude: %v\n", document.Path, document.Latitude, document.Longitude)
		filter := bson.M{"md5": document.Md5}
		count, err := col.CountDocuments(context.TODO(), filter)
		if err != nil {
			fmt.Println("Can not search by Md5")
		}

		if count >= 1 {
			fmt.Printf("File: %v already exist, passed.\n", Red(document.Path))
			continue
		} else {
			counter++
			if counter%30 == 0 {
				fmt.Println("Let me sleep for 90 seconds")
				time.Sleep(90 * time.Second)
			}
			gps := model.ReverseGeocoding(document.Latitude, document.Longitude)
			if len(gps.Address) >= 1 {

				gpsAddress := ""
				for _, v := range gps.Address {
					value := fmt.Sprintf("%v", v)
					gpsAddress = gpsAddress + "," + value
				}

				gpsAddress = gpsAddress[1:len(gpsAddress)]

				fmt.Printf("Address: %v\n", gpsAddress)
				insert := bson.M{"md5": document.Md5, "gpsAddress": gpsAddress}
				_, err := col.InsertOne(context.TODO(), insert)
				if err != nil {
					panic(err)
				}

			} else {
				continue
			}

		}
	}
}

func connectToDB(uri, dbname string, collection string) (*mongo.Collection, error) {

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("mongo.Connect() ERROR:", err)
		os.Exit(1)
	}
	col := client.Database(dbname).Collection(collection)

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
	}
	return col, nil

}
