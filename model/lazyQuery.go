package model

import (
	"context"
	"fmt"

	"github.com/photoServer/global"
	"go.mongodb.org/mongo-driver/mongo"
)

func NextPhoto(cursor *mongo.Cursor) func() (*mongo.Cursor, string) {
	return func() (*mongo.Cursor, string) {
		var document global.Document
		if cursor.Next(context.TODO()) {
			err := cursor.Decode(&document)
			if err != nil {
				fmt.Println("Cursor Decode error in cursor.Next()", err)
				return nil, ""
			}
		}
		photoBase64 := GenOriginalPicBase64(document.Path)
		return cursor, photoBase64

	}
}

// col, err := connectToPic()
// if err != nil {
// 	fmt.Pr*mongo.Cursorln("Error Can not connect to PIC collection")
// 	return "", err
// }
// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// defer cancel()

// Demo

// func f(i int) func() int {
//      return func() int {
//          i++
//          return i
//      }
// }

// 函数f返回了一个函数，返回的这个函数就是一个闭包。这个函数中本身是没有定义变量i的，而是引用了它所在的环境（函数f）中的变量i。

// 我们再看一下效果：

// c1 := f(0)
// c2 := f(0)
// c1() // 打印 1
// c2() // 打印 1
