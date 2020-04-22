package model

import (
	"context"

	"github.com/photoServer/global"
	"go.mongodb.org/mongo-driver/bson"
)

func NextPhoto() func() string {
	//Closure function, search all record
	//Usage:
	// c1 := model.NextPhoto()
	// c1()
	// c1()

	col, _ := ConnectToPic()
	cursor, err := col.Find(context.TODO(), bson.M{})
	if err != nil {
		panic(err)
	}

	return func() string {
		if cursor.Next(context.TODO()) {
			var document global.Document
			err := cursor.Decode(&document)
			if err != nil {
				panic(err)
			}
			base64Pic := GenOriginalPicBase64(document.Path)
			return base64Pic
		} else {
			return ""
		}

	}
}

// func handleRoute() iris.Handler {
// 	c1 := NextPhoto2()
// 	return func(ctx iris.Context) {
// 		ctx.JSON(iris.Map{"status": iris.StatusOK, "message": c1()})

// 	}
// }
