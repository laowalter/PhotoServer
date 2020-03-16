package main

import (
	"fmt"

	"github.com/photoServer/mymongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoFields struct {
	Key         string             `json:"key,omitempty"`
	ID          primitive.ObjectID `bson:"_id, omitempty"`
	StringField string             `bson:"string field" json:"string field"`
	IntField    int                `bson:"int field" json:"int field"`
	BoolField   bool               `bson:"bool field" json:"bool field"`
}

func main() {
	client, ctx := mymongo.ConnectDB("mongodb://localhost:27017")
	db := &mymongo.DBconn{Client: client, Ctx: ctx, DBName: "persons"}
	/*
		document1 := &MongoFields{
			Key:         "record1",
			StringField: "stringfield1",
			IntField:    1,
			BoolField:   true,
		}
	*/
	result := &MongoFields{}
	idStr := "5e6dc7e7ea9741af391067e9"
	docID, err := primitive.ObjectIDFromHex(idStr)

	if err != nil {
		panic(err)
	}
	/*	err = db.Client.Database("persons").Collection("some").FindOne(ctx, bson.M{"_id": docID}).Decode(&result)
	 */
	fmt.Println(result)

	r := db.ReadOne("some", ctx, bson.M{"_id": docID})
	result = r.(*MongoFields)
	fmt.Println(result)

}
