package model

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func ListTags() ([]string, error) {
	//Find all the distinct tags.
	col, err := ConnectToPic()
	if err != nil {
		fmt.Println("Error Can not connect to PIC collection")
		return []string{}, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{}
	result, err := col.Distinct(ctx, "tags", filter) //[]interface{} [101 日本 西三旗]
	if err != nil {
		fmt.Printf("Can not find Distinct of tags")
		return []string{}, err
	}

	tags := make([]string, len(result))
	for i, v := range result {
		tags[i] = fmt.Sprint(v)
	}
	return tags, nil
}
