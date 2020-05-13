package main

import (
	"context"
	"fmt"
	"os"
	"strings"

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
	for _, document := range documentList {
		strings.TrimSpace(document.Latitude)
		strings.TrimSpace(document.Longitude)

		if document.Latitude == "" || document.Longitude == "" {
			continue
		}

		fmt.Printf("Dealing with: %v, latitude: %v, longitude: %v\n", document.Path, document.Latitude, document.Longitude)
		gps := model.ReverseGeocoding(document.Latitude, document.Longitude)
		gpsAddress := ""
		for _, v := range gps.Address {
			value := fmt.Sprintf("%v", v)
			gpsAddress = gpsAddress + "," + value
		}
		gpsAddress = gpsAddress[1:len(gpsAddress)]
		fmt.Printf("Address: %v\n", gpsAddress)
		updateGPSDB(col, document, gpsAddress)
	}
}

func updateGPSDB(collection *mongo.Collection, document global.Document, address string) {
	insert := bson.M{"document.Md5": document.Md5, "gpsAddress": address}
	_, insertErr := collection.InsertOne(context.TODO(), insert)
	if insertErr != nil {
		//check if the Md5 already exist in db is yes
		filter := bson.M{"md5": document.Md5}
		count, err := collection.CountDocuments(context.TODO(), filter)
		if err != nil {
			fmt.Println("Can not search by Md5")
		}
		if count >= 1 {
			update := bson.M{"$set": bson.M{"document.Md5": document.Md5, "gpsAddress": address}}
			_ = collection.FindOneAndUpdate(context.TODO(), filter, update)
			fmt.Printf("File: %v already exist, filename and path updated.\n", Red(document.Path))
			return

		} else {
			insert := bson.M{"document.Md5": document.Md5, "gpsAddress": address}
			_, err := collection.InsertOne(context.TODO(), insert)
			if err != nil {
				panic(err)
			}

		}

	}
	return
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
