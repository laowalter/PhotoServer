package main

import (
	"fmt"
	"runtime"
)

/*
Ref: https://www.socketloop.com/tutorials/golang-generate-thumbnails-from-images
*/

func main() {
	// use all CPU cores for maximum performance
	runtime.GOMAXPROCS(runtime.NumCPU())

	// input files
	path := "/home/walter/go/src/github.com/photoServer/testData"
	files := ListPicVideoFilesTree(path)
	for _, file := range files {
		fmt.Println(file)
	}

	/*
		// load images and make 100x100 thumbnails of them
		var thumbnails []image.Image
		for _, file := range files {
			img, err := imaging.Open(file)
			if err != nil {
				panic(err)
			}
			thumb := imaging.Thumbnail(img, 100, 100, imaging.CatmullRom)
			thumbnails = append(thumbnails, thumb)
		}

		// create a new blank image
		dst := imaging.New(100*len(thumbnails), 100, color.NRGBA{0, 0, 0, 0})

		// paste thumbnails into the new image side by side
		for i, thumb := range thumbnails {
			dst = imaging.Paste(dst, thumb, image.Pt(i*100, 0))
		}

		// save the combined image to file
		err := imaging.Save(dst, "dst.jpg")
		if err != nil {
			panic(err)
		}
	*/
}
