package model

import (
	"context"
	"fmt"

	"github.com/photoServer/global"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
func connectToDB(uri, dbname string) (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("mongo.Connect() ERROR:", err)
		os.Exit(1)
	}
	return client.Database(dbname), nil
}
*/

func ConnectToPic() (*mongo.Collection, error) {
	database, collection, uri := global.DBname, global.PICcol, global.MongoUri
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	var col *mongo.Collection
	if err != nil {
		fmt.Println("mongo.Connect() ERROR:", err)
		return col, err
	}
	col = client.Database(database).Collection(collection)
	return col, nil
}
