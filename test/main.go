package main

import (
	"github.com/photoServer/model"
)

func main() {
	lon := `138 deg 48' 39.84" E`
	lat := `35 deg 9' 7.22" N`
	model.ReverseGeocoding(lat, lon)
}
