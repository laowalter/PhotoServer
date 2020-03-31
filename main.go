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

	tmpl := iris.HTML("./views/templates", ".html")
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

	totalPages := model.CountDocumentsPages() / global.PhotosPerPage
	currentPage := rand.Int63n(totalPages) //can random current page:w

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

//http://192.168.0.199:8080/page?year=1980&page=23
func (c *RootController) GetPage(ctx iris.Context) mvc.Result {
	fmt.Printf("FullURL: %s\n", ctx.URLParam("year"))
	fmt.Printf("FullURL: %s\n", ctx.URLParam("page"))
	fmt.Println("-------------------------")

	return mvc.View{}
}

func (c *RootController) GetBy(year int) mvc.Result {
	//处理localhost:8080/2019的控制器
	yearPic, totalPages, err := model.QueryPhotosByYear(year, int64(1))
	if err != nil {
		fmt.Println("Finding all thumbnail by year")
	}

	//pagers := util.Pagers(currentPage, totalPages)
	return mvc.View{
		Name: "index.html",
		Data: iris.Map{
			"years":       GlobalYearList,
			"thumb":       yearPic,
			"totalpages":  totalPages,
			"currentYear": year,
		},
	}
}

func photoMVC(app *mvc.Application) {
	//用来处理 localhost:8080/photo的MVC
	app.Router.Use(func(ctx iris.Context) {
		ctx.Application().Logger().Infof("Path: %s", ctx.Path())
		ctx.Next()
	})
	app.Handle(new(PhotoController))
}

func (c *PhotoController) GetBy(path string) mvc.Result {
	//显示图片的原始文件
	//通过GetBy获得thumb图片对应的href中的原图对应的全路径(path)
	//这个全路径为了能在url中通过GetBy获得，已经将path(如/data/1.jpg)
	//进行了base64转换。现在需要先decode回来。

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
