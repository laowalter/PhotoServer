package mymongo

/*
Ref https://blog.ruanbekker.com/blog/2019/04/17/mongodb-examples-with-golang/
InsertOne: Insert one record.
InsertMany: Insert multiple record once
*/

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DBconn struct {
	Client *mongo.Client
	Ctx    context.Context
	DBName string //Database Name
}

func ConnectDB(uri string) (*mongo.Client, context.Context) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	return client, ctx
}

func (db DBconn) ListDBNames() []string {
	databases, _ := db.Client.ListDatabaseNames(db.Ctx, bson.M{})
	return databases
}

func (db DBconn) InsertOne(collectionName string, data interface{}) error {
	collection := db.Client.Database(db.DBName).Collection(collectionName)
	//insertResult, err := collection.InsertOne(context.TODO(), data)
	_, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		return err
	}
	//fmt.Println("Inserted a Single Document: ", insertResult.InsertedID)
	return nil
}

func (db DBconn) InsertMany(collectionName string, data []interface{}) error {
	collection := db.Client.Database(db.DBName).Collection(collectionName)
	//InsertManyResult, err := collection.InsertMany(context.TODO(), data)
	_, err := collection.InsertMany(context.TODO(), data)
	if err != nil {
		return err
	}
	//fmt.Println("Inserted a Single Document: ", InsertManyResult.InsertedIDs)
	return nil
}

func (db DBconn) ReadOne(collectionName string, ctx context.Context, filter interface{}) interface{} {
	return db.Client.Database(db.DBName).Collection(collectionName).FindOne(ctx, filter)
}

/*
package main

import (
        "github.com/photoServer/mymongo"
)

type Person struct {
        Name string
        Age  int
        City string
}

func main() {
        client, ctx := mymongo.ConnectDB("mongodb://localhost:27017") //Just connect monogodb
        db := &mymongo.DBconn{Client: client, Ctx: ctx, DBName: "persons"}

        err := db.ReadOne("mydb")
        if err != nil {
                panic(err)
        }

}
*/
