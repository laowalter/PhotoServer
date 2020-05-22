package main

/* Generate thumbnail and store to mongodb
Database: album
Collection: pic
*/
import (
	"context"
	"fmt"
	"os"

	. "github.com/logrusorgru/aurora"
	"github.com/photoServer/global"
	"github.com/photoServer/model"
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
	cleanPic()
	cleanGps()
}

func cleanGps() {
	colGps, err := model.ConnectToGps()
	if err != nil {
		fmt.Println("Can not open Collection gps")
	}

	colPic, err := model.ConnectToPic()
	if err != nil {
		fmt.Println("Can not open Collection gps")
	}

	cursor, err := colGps.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println("Can not find data in Collection gps")
		return
	} else {
		var gpsResult struct {
			Md5        string `bson:"md5"`
			GpsAddress string `bson:"gpsAddress"`
		}

		for cursor.Next(context.TODO()) {
			err := cursor.Decode(&gpsResult)
			if err != nil {
				fmt.Println("cursor.Next() error:", err)
				os.Exit(1)
			}

			//fmt.Println("Md5", gpsResult.Md5)
			filter := bson.M{"md5": gpsResult.Md5}

			count, err := colPic.CountDocuments(context.TODO(), filter)
			if err != nil {
				fmt.Println(err)
				fmt.Println("Can not count document in collection pic")
				continue
			}

			if count == 0 {
				_, err = colGps.DeleteOne(context.TODO(), filter)
				if err != nil {
					fmt.Println(err)
					fmt.Printf("Can not remove the Record with Md5: %v\n", Red(gpsResult.Md5))
					continue
				}
				fmt.Printf("md5: %v removed from collection gps\n", Red(gpsResult.Md5))
			}
		}
	}
}

func cleanPic() {
	database, collection, uri := "album", "pic", "mongodb://localhost:27017"
	db, _ := connectToDB(uri, database)
	col := db.Collection(collection)
	count, err := col.CountDocuments(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println("Collection counting error!")
		return
	}

	fmt.Println("The total files in database: ", Green(count))

	removedNum := 0

	cursor, err := col.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println("Finding all documents ERROR:", err)
		defer cursor.Close(context.TODO())
	} else {
		for cursor.Next(context.TODO()) {
			//fmt.Println(elem[0])
			var result global.Document
			err := cursor.Decode(&result)
			if err != nil {
				fmt.Println("cursor.Next() error:", err)
				os.Exit(1)
			}

			_, err = os.Stat(result.Path)
			if os.IsNotExist(err) {
				fmt.Printf("%s  %v does not exist", Red("OOPs!"), Red(result.Path))
				_, err := col.DeleteOne(context.TODO(), bson.M{"_id": cursor.Current.Lookup("_id")})
				if err != nil {
					fmt.Println(err)
					fmt.Printf("%s can not remove record %s from database.", Red("Oops!"), Red(result.Path))
				} else {
					removedNum++
					fmt.Printf(" %s from database!\n", Red("removed"))
					//fmt.Println(delresult)
				}
			}
		}
		cursor.Close(context.TODO())
	}
	fmt.Printf("%v files removed from database.", Red(removedNum))
	fmt.Println("All Document Cleaned!")
}
