package model

import (
	"context"
	"fmt"
	"time"

	"github.com/photoServer/global"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetYearList() ([]*global.YearCount, error) {
	var yearCount []*global.YearCount
	col, err := ConnectToPic()
	if err != nil {
		fmt.Println("Error Can not connect to PIC collection")
		return yearCount, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Second)
	defer cancel()

	/* The following pipline:
	db.pic.aggregate({$group:
		{ _id:   {year:{$year:"$createdate"}},
		  counter:{$sum:1}
	    }},
		{$sort:{"_id.year":-1}}
	  )
	*/
	pipeline := []bson.M{bson.M{"$group": bson.M{"_id": bson.M{"year": bson.M{"$year": "$createdate"}},
		"counter": bson.M{"$sum": 1},
	}},
		bson.M{"$sort": bson.M{"_id.year": -1}},
	}

	cursor, err := col.Aggregate(ctx, pipeline)
	if err != nil {
		fmt.Println("Aggregate Error", err)
		return yearCount, err

	}
	defer cursor.Close(ctx)
	/*
		var result []bson.M
		cursor.All(ctx, &result)
		fmt.Println(result)
		//output: map[_id:map[year:2019] counter:1]
	*/

	for cursor.Next(ctx) {
		var _result bson.M
		if err := cursor.Decode(&_result); err != nil {
			fmt.Println("Can not decode Aggregate result")
		}
		_year := _result["_id"].(primitive.M)
		year := _year["year"].(int32)
		number := _result["counter"].(int32)
		data := &global.YearCount{
			Year:   year,
			Number: number,
		}
		//fmt.Println(data)
		yearCount = append(yearCount, data)
	}
	return yearCount, nil
}
