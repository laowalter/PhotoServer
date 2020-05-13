package rootview

import (
	"fmt"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/photoServer/model"
)

func SinglePhotoView(md5 string) mvc.Result {

	document, err := model.QueryPhotoByMd5(md5)

	if err != nil {
		fmt.Println("Can to get the Single document by md5.")
	}

	photoBase64 := model.GenOriginalPicBase64(document.Path)
	// fmt.Println("document.md5: ", document.Md5)
	// fmt.Println("md5: ", md5)

	gps, err := model.QueryGPSByMd5(document.Md5)
	if err != nil {
		fmt.Println("model.QueryGPS query error")
	}

	return &mvc.View{
		Name: "singlePhoto.html",
		Data: iris.Map{
			"thumb": photoBase64,
			"photo": document,
			"gps":   gps,
		},
	}

}
