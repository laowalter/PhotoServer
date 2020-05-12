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

	return &mvc.View{
		Name: "singlePhoto.html",
		Data: iris.Map{
			"thumb": photoBase64,
			"photo": document,
		},
	}

}
