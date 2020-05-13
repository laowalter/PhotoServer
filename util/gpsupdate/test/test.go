package main

import (
	"fmt"

	"github.com/photoServer/model"
)

func main() {
	lat := `22 deg 31' 29.40" N`
	lon := `113 deg 59' 6.60" E`

	gps := model.ReverseGeocoding(lat, lon)
	gpsAddress := ""

	for _, v := range gps.Address {
		value := fmt.Sprintf("%v", v)
		gpsAddress = gpsAddress + "," + value
	}
	gpsAddress = gpsAddress[1:len(gpsAddress)]
	fmt.Printf("Address: %v\n", gpsAddress)

}
