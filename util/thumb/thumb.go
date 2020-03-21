package main

import (
	"fmt"
	"os"

	"github.com/disintegration/imaging"
)

func main() {
	file := os.Args[1]
	img, _ := imaging.Open(file, imaging.AutoOrientation(true))
	rectangle := img.Bounds()

	fmt.Println(rectangle.Dx(), rectangle.Dy())
	t_height := 250
	t_width := int(rectangle.Dx() * 250 / rectangle.Dy())
	fmt.Println(t_width, t_height)
}
