package main

import (
	"fmt"
	"math/rand"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/recover"
	"github.com/photoServer/global"
	"github.com/photoServer/model"
	"github.com/photoServer/view/rootview"

	"github.com/kataras/iris/mvc"
)

var GlobalYearList []*global.YearCount

func main() {
	app := iris.New()
	app.Use(recover.New())
	app.Logger().SetLevel("debug")

	app.HandleDir("/static", "./assets")
	app.Favicon("./assets/favicon.ico")

	tmpl := iris.HTML("./view/templates", ".html")
	app.RegisterView(tmpl)

	mvc.Configure(app.Party("/"), rootMVC)
	app.Run(iris.Addr(":8080"))
}

type RootController struct{}

func rootMVC(app *mvc.Application) {
	app.Router.Use(func(ctx iris.Context) {
		ctx.Application().Logger().Infof("Path: %s", ctx.Path())
		ctx.Next()
	})

	app.Handle(new(RootController))
}

//http://192.168.0.199:8080/
//http://192.168.0.199:8080/?page=123
func (c *RootController) Get(ctx iris.Context) mvc.Result {

	if ctx.IsMobile() {
		fmt.Println("from mobile dev")
	} else {
		fmt.Println("from desktop")
	}

	currentPage, err := ctx.URLParamInt64("page")
	if err != nil {
		totalPages := model.CountDocumentsPages() / global.PhotosPerPage
		currentPage = rand.Int63n(totalPages) //can random current page
	}

	view := rootview.RootView(currentPage)
	return view
}

//http://192.168.0.199:8080/year?year=2019page=23
func (c *RootController) GetYear(ctx iris.Context) mvc.Result {
	currentYear, err := ctx.URLParamInt("year")
	if err != nil {
		fmt.Println("Did not get year")
	}
	currentPage, err := ctx.URLParamInt64("page")
	if err != nil {
		fmt.Println("Did not currentPage")
	}

	view := rootview.YearView(currentYear, currentPage)
	return view
}

//http://192.168.0.199:8080/single?md5=md5string
func (c *RootController) GetSingle(ctx iris.Context) mvc.Result {
	md5 := ctx.URLParam("md5")
	view := rootview.SinglePhotoView(md5)
	return view
}
