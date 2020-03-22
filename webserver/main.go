package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"

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

	documentList := getPicList(col)

	r := gin.Default()
	r.Static("/photo", "./statics")

	r.SetFuncMap(template.FuncMap{
		"saft": func(str string) template.HTML {
			return template.HTML(str)
		},
	})

	r.LoadHTMLGlob("templates/*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "./index.html", gin.H{
			"piclist": documentList,
		})
	})

	r.Run("192.168.0.199:4000")
}

func getPicList(col *mongo.Collection) []global.Document {

	var documentList []global.Document
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	cursor, err := col.Find(ctx, bson.D{})
	if err != nil {
		fmt.Println("Finding all documents ERROR:", err)
		defer cursor.Close(ctx)
		return documentList
	} else {
		for cursor.Next(ctx) {
			var document global.Document
			err := cursor.Decode(&document)
			if err != nil {
				fmt.Println("cursor.Next() error:", err)
				os.Exit(1)
			}
			//document.Thumbase64 = base64.StdEncoding.EncodeToString(document.Thumbnail)
			documentList = append(documentList, document)
		}
	}

	return documentList
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