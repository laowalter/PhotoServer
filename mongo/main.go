package main

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	database, collection, uri := "album", "pic", "mongodb://localhost:27017"
	db, _ := connectToDB(uri, database)
	col := db.Collection(collection)
	getYearList(col)
}

func getYearList(col *mongo.Collection) {
	//ctx, cancel := context.WithTimeout(context.Background(), 200*time.Second)
	//defer cancel()
	ctx := context.TODO()
	/*
		db.pic.aggregate({$group:
			{ _id:   {year:{$year:"$createtime"}},
			  counter:{$sum:1}
		    }},
			{$sort:{"_id.year":-1}}
		  )
	*/
	pipeline := []bson.M{bson.M{"$group": bson.M{"_id": bson.M{"year": bson.M{"$year": "$createtime"}},
		"counter": bson.M{"$sum": 1},
	}},
		bson.M{"$sort": bson.M{"_id.year": -1}},
	}

	cursor, err := col.Aggregate(ctx, pipeline)
	if err != nil {
		fmt.Println("Aggregate Error", err)
		return

	}
	defer cursor.Close(ctx)

	/*
		var result []bson.M
		cursor.All(ctx, &result)
		fmt.Println(result)
		//output: map[_id:map[year:2019] counter:1]
	*/
	for cursor.Next(ctx) {
		var result bson.M

		if err := cursor.Decode(&result); err != nil {
			fmt.Println("Can not decode Aggregate result")
		}
		fmt.Printf("%v, %T\n", result["_id"], result["_id"])
		year := result["_id"].(primitive.M)
		fmt.Println(year["year"])
	}

	return
}

func connectToDB(uri, dbname string) (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("mongo.Connect() ERROR:", err)
		os.Exit(1)
	}
	return client.Database(dbname), nil
}
