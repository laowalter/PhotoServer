package main

import (
	"context"
	"fmt"
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

func (db DBconn) listDBNames() []string {
	databases, _ := db.Client.ListDatabaseNames(db.Ctx, bson.M{})
	return databases
}

func (db DBconn) insertDB(collectionName string, data interface{}) error {
	collection := db.Client.Database(db.DBName).Collection(collectionName)
	insertResult, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("Inserted a Single Document: ", insertResult.InsertedID)
	return nil
}

func connectDB(uri string) (*mongo.Client, context.Context) {
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

type Person struct {
	Name string
	Age  int
	City string
}

func main() {
	client, ctx := connectDB("mongodb://localhost:27017") //Just connect monogodb
	db := &DBconn{Client: client, Ctx: ctx, DBName: "persons"}
	dbNames := db.listDBNames()
	fmt.Println(dbNames)
	ruan := Person{"Ruan", 34, "Cape Town"}
	err := db.insertDB("mydb", ruan)
	if err != nil {
		panic(err)
	}
}
