package model

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/photoServer/global"
	"go.mongodb.org/mongo-driver/bson"
)

func DeletePhotos(filePathList []string) {
	var removed []string
	for _, oldFile := range filePathList {
		_, newFile := filepath.Split(oldFile)
		newFile = global.RemovedDir + "/" + newFile
		err := os.Rename(oldFile, newFile)
		if err != nil {
			fmt.Printf("File: %s can not move to %s\n", oldFile, global.RemovedDir)
			fmt.Println(err)
			continue
		}
		removed = append(removed, oldFile)
		//fmt.Println(newFile, " moved")

		removeThumbFromDB(removed)
	}

}

func removeThumbFromDB(removed []string) {
	col, err := connectToPic()
	if err != nil {
		fmt.Println("Error Can not connect to PIC collection")
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for _, file := range removed {
		filter := bson.M{"path": file}
		_, err := col.DeleteOne(ctx, filter)
		if err != nil {
			fmt.Println(err)
			fmt.Printf("%s can not remove record from database.", file)
			continue
		}

		//fmt.Printf("Photo: %s removed from DB.\n", file)
	}
}
