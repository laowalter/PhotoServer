package main

import (
	"fmt"
	"strings"

	//. "github.com/logrusorgru/aurora"
	"github.com/photoServer/model"
)

func main() {

	documentList, err := model.QueryAllPhotosOnepage()
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
		for k, v := range gps.Address {
			fmt.Println(k, v)
		}

	}

	// database, collection, uri := "album", "gps", "mongodb://localhost:27017"
	// db, _ := connectToDB(uri, database)
	// col := db.Collection(collection)

	// { //Create index by md5
	// 	mod := mongo.IndexModel{
	// 		Keys: bson.M{
	// 			"md5": 1,
	// 		}, Options: options.Index().SetUnique(true),
	// 	}
	// 	_, err := col.Indexes().CreateOne(context.TODO(), mod)
	// 	if err != nil { // Check if the CreateOne() method returned any errors
	// 		fmt.Println("Indexes().CreateOne() ERROR:", err)
	// 		os.Exit(1)
	// 	} else { // API call returns string of the index name
	// 		fmt.Printf("Notice: Use %s of file as the unique index.\n", Red("md5sum"))
	// 	}
	// }

	// _, insertErr := col.InsertOne(context.TODO(), document)

}

// func connectToDB(uri, dbname string) (*mongo.Database, error) {

// 	clientOptions := options.Client().ApplyURI(uri)
// 	client, err := mongo.Connect(context.TODO(), clientOptions)
// 	if err != nil {
// 		fmt.Println("mongo.Connect() ERROR:", err)
// 		os.Exit(1)
// 	}
// 	return client.Database(dbname), nil
// }
