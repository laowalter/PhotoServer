package main

import (
	"github.com/photoServer/mymongo"
)

type Person struct {
	Name string
	Age  int
	City string
}

func main() {
	client, ctx := mymongo.ConnectDB("mongodb://localhost:27017") //Just connect monogodb
	db := &mymongo.DBconn{Client: client, Ctx: ctx, DBName: "persons"}

	err := db.ReadOne("mydb")
	if err != nil {
		panic(err)
	}

}
