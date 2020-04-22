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
}

func SlideAnyView() mvc.Result {
	return &mvc.View{
		Name: "slideany.html",
		Data: iris.Map{},
	}
}
