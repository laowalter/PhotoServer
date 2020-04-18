package model

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func AddTags(tagsWithFiles string) {
	_tmp := strings.Split(tagsWithFiles, "|") // "tags|filePathList"
	_tags, _files := _tmp[0], _tmp[1]
	_tags = strings.TrimSpace(_tags)

	space := regexp.MustCompile(`\s+`)
	_tags = space.ReplaceAllString(_tags, " ")
	tags := strings.Split(_tags, " ")
	files := strings.Split(_files, ",")
	col, err := ConnectToPic()

	if err != nil {
		fmt.Println("Error Can not connect to PIC collection")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for _, file := range files {
		filter := bson.M{"path": file}
		fmt.Println(file)
		fmt.Println(tags)
		updater := bson.M{
			"$addToSet": bson.M{
				"tags": bson.M{"$each": tags},
			},
		}
		res, err := col.UpdateOne(ctx, filter, updater)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v", res)
	}
}
