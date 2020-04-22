package main

import (
	"context"
	"fmt"

	"github.com/photoServer/global"
	"github.com/photoServer/model"
	"go.mongodb.org/mongo-driver/bson"
)

func NextPhoto2() func() string {
	//Closure function, search all record
	//Usage:
	// c1 := model.NextPhoto()
	// c1()
	// c1()

	col, _ := model.ConnectToPic()
	cursor, err := col.Find(context.TODO(), bson.M{})
	if err != nil {
		panic(err)
	}

	return func() string {
		if cursor.Next(context.TODO()) {
		} else {
			return ""
		}
		var document global.Document
		err := cursor.Decode(&document)
		if err != nil {
			panic(err)
		}
		return document.Path
	}
}

func main() {
	c1 := NextPhoto2()
	a := c1()
	fmt.Println(a)

	a = c1()
	fmt.Println(a)

	a = c1()
	fmt.Println(a)

	a = c1()
	fmt.Println(a)
}
