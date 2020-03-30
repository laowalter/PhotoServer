package model

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"time"

	"github.com/photoServer/global"
	"go.mongodb.org/mongo-driver/bson"
)

func GetThumbByYear(year int) ([]global.Document, error) {
	database, collection, uri := "album", "pic", "mongodb://localhost:27017"
	db, _ := connectToDB(uri, database)
	col := db.Collection(collection)

	var documentList []global.Document
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Second)
	defer cancel()

	fromDate := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
	toDate := time.Date(year+1, time.January, 1, 0, 0, 0, 0, time.UTC)

	filter := bson.M{
		"createtime": bson.M{
			"$gt": fromDate,
			"$lt": toDate,
		},
	}

	cursor, err := col.Find(ctx, filter)
	if err != nil {
		fmt.Println("Finding all documents ERROR:", err)
		defer cursor.Close(ctx)
		return documentList, err
	} else {
		for cursor.Next(ctx) {
			var document global.Document
			err := cursor.Decode(&document)
			if err != nil {
				fmt.Println("cursor.Next() error:", err)
				os.Exit(1)
			}

			//Encode pic path with base64 for iris url param get()
			document.Path = base64.StdEncoding.EncodeToString([]byte(document.Path))
			documentList = append(documentList, document)
		}
	}

	return documentList, nil
}
