package main

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
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

	//db.pic.aggregate({$group:{_id:{year:{$year:"$createtime"}}, counter:{$sum:1}}}, {$sort:{"_id.year":-1}})

	/*
		db.pic.aggregate({$group:
			{ _id:   {year:{$year:"$createtime"}},
			  counter:{$sum:1}
		    }},
			{$sort:{"_id.year":-1}}
		  )
	*/
	//var result []bson.M
	/* OK:
	pipeline := []bson.M{bson.M{"$group": bson.M{"_id": bson.M{"year": bson.M{"$year": "$createtime"}},
		"counter": bson.M{"$sum": -1},
	}}}
	*/

	pipeline := []bson.M{bson.M{"$group": bson.M{"_id": bson.M{"year": bson.M{"$year": "$createtime"}},
		"counter": bson.M{"$sum": -1},
	}},
		bson.M{"$sort": bson.M{"_id.year": -1}},
	}

	var result []bson.M
	//https://docs.mongodb.com/manual/reference/method/db.collection.aggregate/_
	cursor, err := col.Aggregate(ctx, pipeline)
	if err != nil {
		fmt.Println("Aggregate Error", err)
		//defer cursor.Close(ctx)
		return

	} else {

		cursor.All(ctx, &result)
		fmt.Println(result)
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
