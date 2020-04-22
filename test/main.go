package main

import (
	"context"
	"fmt"

	"github.com/kataras/iris"
	"github.com/photoServer/global"
	"github.com/photoServer/model"
	"gopkg.in/mgo.v2/bson"
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

func handleRoute() iris.Handler {
	c1 := NextPhoto2()
	return func(ctx iris.Context) {
		ctx.JSON(iris.Map{"status": iris.StatusOK, "message": c1()})

	}
}

func main() {
	col, _ := model.ConnectToPic()
	cursor, err := col.Find(context.TODO(), bson.M{})
	if err != nil {
		panic(err)
	}

	var i = 0
	for {

		if cursor.Next(context.TODO()) {
			fmt.Println("Number:", i)
			i++
		} else {
			fmt.Println("No cursor NEXT()")
			break
		}
	}
}
