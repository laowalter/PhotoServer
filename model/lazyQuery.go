package model

import (
	"context"
	"fmt"

	"github.com/photoServer/global"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NextPhoto() func() *mongo.Cursor {
	//Closure function, search all record, return one fro every apply
	//Usage:
	// c1 := model.NextPhoto()
	// c1()
	// c1()

	col, _ := ConnectToPic()
	cursor, _ := col.Find(context.TODO(), bson.M{})

	return func() *mongo.Cursor {

		if cursor.Next(context.TODO()) {
			var document global.Document
			err := cursor.Decode(&document)
			if err != nil {
				return nil
			}
			fmt.Println(document.Path)
		}
		return cursor
	}
}
