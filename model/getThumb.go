package model

import (
	"context"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/photoServer/global"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetThumbByYear(year int, pageNumber int64) ([]global.Document, error) {

	fromDate := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
	toDate := time.Date(year+1, time.January, 1, 0, 0, 0, 0, time.UTC)

	filter := bson.M{
		"createtime": bson.M{
			"$gt": fromDate,
			"$lt": toDate,
		},
	}
	var documentList []global.Document
	documentList, err := getDocument(filter, pageNumber)
	if err != nil {
		return documentList, err
	}
	return documentList, nil
}

func GetThumbList(pageNumber int64) ([]global.Document, error) {
	var documentList []global.Document
	filter := bson.M{}
	documentList, err := getDocument(filter, pageNumber)
	if err != nil {
		return documentList, err
	}
	return documentList, nil

}

func getDocument(filter bson.M, pageNumber int64) ([]global.Document, error) {
	database, collection, uri := "album", "pic", "mongodb://localhost:27017"
	db, _ := connectToDB(uri, database)
	col := db.Collection(collection)

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Second)
	defer cancel()

	var documentList []global.Document

	var opt options.FindOptions
	cursor, err := col.Find(
		ctx,
		filter,
		opt.SetSkip((pageNumber-1)*global.PhotosPerPage),
		opt.SetLimit(global.PhotosPerPage),
	)
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
				return documentList, err
			}

			//Encode pic path with base64 for iris url param get()
			document.Path = base64.StdEncoding.EncodeToString([]byte(document.Path))
			documentList = append(documentList, document)
		}
	}

	return documentList, nil
}
