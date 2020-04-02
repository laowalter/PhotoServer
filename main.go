package main

import (
	"encoding/base64"
	"fmt"
	"math/rand"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/recover"
	"github.com/photoServer/global"
	"github.com/photoServer/model"
	"github.com/photoServer/views/util"

	"github.com/kataras/iris/mvc"
)

var GlobalYearList []*global.YearCount

func main() {
	app := iris.New()
	app.Use(recover.New())
	app.Logger().SetLevel("debug")

	app.HandleDir("/static", "./assets")
	app.Favicon("./assets/favicon.ico")

	tmpl := iris.HTML("./views/templates", ".html")
	app.RegisterView(tmpl)

	mvc.Configure(app.Party("/"), rootMVC)

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

//http://192.168.0.199:8080/
//http://192.168.0.199:8080/?page=123
func (c *RootController) Get(ctx iris.Context) mvc.Result {

	currentPage, err := ctx.URLParamInt64("page")
	if err != nil {
		totalPages := model.CountDocumentsPages() / global.PhotosPerPage
		currentPage = rand.Int63n(totalPages) //can random current page
	}

	YearList, err := model.GetYearList()
	GlobalYearList = YearList
	if err != nil {
		fmt.Println("Can not Get Year List")
	}

	picList, totalPages, err := model.QueryAllPhotos(currentPage)
	if err != nil {
		fmt.Println("Can not Get Thumbnail List")
	}

	pagers := util.Pagers(currentPage, totalPages)
	return mvc.View{
		Name: "index.html",
		Data: iris.Map{
			"years":      YearList,
			"thumb":      picList,
			"totalpages": totalPages,
			"pagers":     pagers,
		},
	}
}

//http://192.168.0.199:8080/year?year=2019page=23
func (c *RootController) GetYear(ctx iris.Context) mvc.Result {
	year, err := ctx.URLParamInt("year")
	if err != nil {
		fmt.Println("Did not get year")
	}
	currentPage, err := ctx.URLParamInt64("page")
	if err != nil {
		fmt.Println("Did not currentPage")
	}

	yearPic, totalPages, err := model.QueryPhotosByYear(year, currentPage)
	if err != nil {
		fmt.Println("Can not finding all thumbnail by year")
	}

	pagers := util.Pagers(currentPage, totalPages)
	return mvc.View{
		Name: "index.html",
		Data: iris.Map{
			"years":       GlobalYearList,
			"thumb":       yearPic,
			"totalpages":  totalPages,
			"currentyear": year,
			"pagers":      pagers,
		},
	}
}

//http://192.168.0.199:8080/large?path=base64Enoded_real_path&md5=md5string
func (c *RootController) GetLarge(ctx iris.Context) mvc.Result {
	md5 := ctx.URLParam("md5")
	path := ctx.URLParam("path")
	pathByte, err := base64.StdEncoding.DecodeString(path)
	if err != nil {
		fmt.Println("Can to Decode  \"Path\" from base64.")
	}
	path = string(pathByte)
	OriginalPic := model.GenOriginalPicBase64(path)
	document, err := model.QueryPhotoByMd5(md5)
	if err != nil {
		fmt.Println("Can to get the Single document by md5.")
	}

	return mvc.View{
		Name: "originalPic.html",
		Data: iris.Map{
			"thumb": OriginalPic,
			"photo": document,
		},
	}

}
