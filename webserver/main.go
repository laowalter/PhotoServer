package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/photoServer/global"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	database, collection, uri := "album", "pic", "mongodb://localhost:27017"
	db, _ := connectToDB(uri, database)
	col := db.Collection(collection)

	result := &global.Document{}
	filter := bson.M{"filename": bson.M{"$eq": "DSC00181.JPG.jpg"}}
	err := col.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.Thumbnail)

	thumb := base64.StdEncoding.EncodeToString(result.Thumbnail)
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "./index.html", gin.H{
			"title": "www.eliglad.com",
			"thumb": thumb,
		})
	})

	r.Run("192.168.0.199:4000")
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
