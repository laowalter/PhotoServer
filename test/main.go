package main

import (
	"fmt"

	"github.com/photoServer/model"
)

func main() {
	md5 := "0943dabac90c2bb7fad07fe039c4110f"
	md5 = "a5965d185b3a05ca889cdf0782f38af4"
	gps, err := model.QueryGPSByMd5(md5)
	if err != nil {
		panic(err)
	}

	fmt.Println(gps)

}
