package main

import (
	"encoding/base64"
	"fmt"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/recover"
	"github.com/photoServer/global"
	"github.com/photoServer/model"

	"github.com/kataras/iris/mvc"
)

var GlobalYearList []*global.YearCount

func main() {
	app := iris.New()
	app.Use(recover.New())
	app.Logger().SetLevel("debug")

	app.HandleDir("/static", "./assets")

	tmpl := iris.HTML("./views", ".html")
	app.RegisterView(tmpl)

	mvc.Configure(app.Party("/"), rootMVC)
	mvc.Configure(app.Party("/photo"), photoMVC)

	app.Run(iris.Addr(":8080"))
}

type RootController struct{}

type PhotoController struct{}

func rootMVC(app *mvc.Application) {
	app.Router.Use(func(ctx iris.Context) {
		ctx.Application().Logger().Infof("Path: %s", ctx.Path())
		ctx.Next()
	})

	app.Handle(new(RootController))
}

func (c *RootController) Get() mvc.Result {
	YearList, err := model.GetYearList()
	GlobalYearList = YearList

	if err != nil {
		fmt.Println("Can not Get Year List")
	}

	PicList, totalPages, err := model.GetThumbList(int64(1))
	if err != nil {
		fmt.Println("Can not Get Thumbnail List")
	}

	return mvc.View{
		Name: "index.html",
		Data: iris.Map{"years": YearList, "thumb": PicList, "totalpages": totalPages},
	}
}

func (c *RootController) GetBy(year int) mvc.Result {
	yearPic, totalPages, err := model.GetThumbByYear(year, int64(1))
	if err != nil {
		fmt.Println("Finding all thumbnail by year")
	}
	/*
		YearList, err := model.GetYearList()
		if err != nil {
			fmt.Println("Can not Get Year List")
		}
	*/
	return mvc.View{
		Name: "index.html",
		Data: iris.Map{"years": GlobalYearList, "thumb": yearPic, "totalpages": totalPages},
	}
}

func photoMVC(app *mvc.Application) {
	app.Router.Use(func(ctx iris.Context) {
		ctx.Application().Logger().Infof("Path: %s", ctx.Path())
		ctx.Next()
	})
	app.Handle(new(PhotoController))
}

func (c *PhotoController) GetBy(path string) mvc.Result { //http://192.168.0.199:8080/photo/1980
	pathByte, err := base64.StdEncoding.DecodeString(path)
	if err != nil {
		panic(err)
	}
	path = string(pathByte)
	OriginalPic := model.GenOriginalPicBase64(path)

	return mvc.View{
		Name: "originalPic.html",
		Data: iris.Map{"thumb": OriginalPic},
	}
}
