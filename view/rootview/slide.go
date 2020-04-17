package rootview

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func SlideView() mvc.Result {

	return &mvc.View{
		Name: "slide.html",
		Data: iris.Map{},
	}
	/*
		photoBase64 := model.GenOriginalPicBase64(Path)
		return &mvc.View{
			Name: "singlePhoto.html",
			Data: iris.Map{
				"thumb": photoBase64,
				"photo": document,
			},
		}
	*/

}
