package main

import (
	"context"
	"fmt"

	"github.com/photoServer/global"
	"github.com/photoServer/model"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	col, _ := model.ConnectToPic()
	cursor, _ := col.Find(context.TODO(), bson.M{})

	if cursor.Next(context.TODO()) {
		var document global.Document
		err := cursor.Decode(&document)
		if err != nil {
			panic(err)
		}
		fmt.Println(document.Path)
	}

	c1 := model.NextPhoto(cursor)
	c2 := model.NextPhoto(cursor)
	c3 := model.NextPhoto(cursor)
	c1()
	c2()
	c3()
}
