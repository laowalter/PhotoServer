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

func QueryPhotosByYear(year int, pageNumber int64) ([]global.Document, int64, error) {
	//分页查询
	//搜索条件: 按年查询
	//返回第pageNumber页面中的对应的document，文档的总数

	fromDate := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
	toDate := time.Date(year+1, time.January, 1, 0, 0, 0, 0, time.UTC)

	filter := bson.M{
		"createtime": bson.M{
			"$gt": fromDate,
			"$lt": toDate,
		},
	}
	var documentList []global.Document
	documentList, totalPages, err := getDocument(filter, pageNumber)
	if err != nil {
		return documentList, 0, err
	}
	return documentList, totalPages, nil
}

func QueryAllPhotos(pageNumber int64) ([]global.Document, int64, error) {
	//分页查询
	//搜索条件: 全部数据
	//返回第pageNumber页面中的对应的document，文档的总数
	var documentList []global.Document
	filter := bson.M{}
	documentList, totalPages, err := getDocument(filter, pageNumber)
	if err != nil {
		return documentList, 0, err
	}
	return documentList, totalPages, nil

}

func CountDocumentsPages() int64 {
	var totalPages int64
	filter := bson.M{}
	database, collection, uri := "album", "pic", "mongodb://localhost:27017"
	db, _ := connectToDB(uri, database)
	col := db.Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Second)
	defer cancel()
	totalNumbers, err := col.CountDocuments(ctx, filter)
	if err != nil {
		fmt.Println("CountDocument() error")
		return 0
	}
	if n := totalNumbers / global.PhotosPerPage; n == 1 {
		totalPages = int64(n) + 1
	} else {
		totalPages = int64(n)
	}
	return totalPages
}

func getDocument(filter bson.M, pageNumber int64) ([]global.Document, int64, error) {
	//分页查询：
	//按照filter查询数据库，将查询结果按照 global包定义的全局变量global.PhotosPerPage
	//进行分页，返回的是第pageNumber页面中的doucment数组，这个filter的总数，这个总数
	//用于确定在templates中的pagination计算。

	var documentList []global.Document
	var totalPages int64

	database, collection, uri := "album", "pic", "mongodb://localhost:27017"
	db, _ := connectToDB(uri, database)
	col := db.Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Second)
	defer cancel()

	totalNumbers, err := col.CountDocuments(ctx, filter)
	if n := totalNumbers / global.PhotosPerPage; n == 1 {
		totalPages = int64(n) + 1
	} else {
		totalPages = int64(n)
	}

	fmt.Printf("Total Documents Number: %v", totalNumbers)
	if err != nil {
		fmt.Println("Database Problem")
		return documentList, totalPages, err
	}

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
		return documentList, totalPages, err

	} else {
		for cursor.Next(ctx) {
			var document global.Document
			err := cursor.Decode(&document)
			if err != nil {
				fmt.Println("cursor.Next() error:", err)
				return documentList, pageNumber, err
			}

			//Encode pic path with base64 for iris url param get()
			document.PathBase64 = base64.StdEncoding.EncodeToString([]byte(document.Path))
			documentList = append(documentList, document)
		}
	}

	return documentList, totalPages, nil
}
