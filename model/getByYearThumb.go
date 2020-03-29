package model

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/photoServer/global"
	"go.mongodb.org/mongo-driver/bson"
)

func GetYearThumbList() ([]global.Document, error) {
	database, collection, uri := "album", "pic", "mongodb://localhost:27017"
	db, _ := connectToDB(uri, database)
	col := db.Collection(collection)

	var documentList []global.Document
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Second)
	defer cancel()

	cursor, err := col.Find(ctx, bson.D{})
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
			//处理全路径为/data/album/ => /适应 r.static(photo)
			document.Path = document.Path[len("/data/album")+1:]
			documentList = append(documentList, document)
		}
	}

	return documentList, nil
}
