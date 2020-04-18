package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"

	"github.com/kataras/iris"
	"github.com/photoServer/model"
	"github.com/photoServer/view/rootview"

	"github.com/kataras/iris/mvc"
)

var TAG = true // if need to initial NextPhoto

func main() {
	app := iris.New()
	//app.Use(recover.New())
	//app.Logger().SetLevel("debug")
	app.Logger().SetLevel("info")

	app.StaticWeb("/static", "./assets")
	app.Favicon("./assets/favicon.ico")

	tmpl := iris.HTML("./assets/templates", ".html")
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

	currentPage, err := ctx.URLParamInt64("page")
	fmt.Println(currentPage)
	if err != nil {
		totalPages, err := model.CountDocumentsPages()
		if err != nil { //Database total photopage calculate wrong
			panic(err)
		}

		currentPage = rand.Int63n(totalPages) //can random current page
		if currentPage == 0 {                 //minmum rand number can be zero.
			currentPage = 1
		}
	}
	view := rootview.RootView(currentPage, ctx.IsMobile())
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
	view := rootview.YearView(currentYear, currentPage, ctx.IsMobile())
	return view
}

//http://192.168.0.199:8080/single?md5=md5string
func (c *RootController) GetSingle(ctx iris.Context) mvc.Result {
	md5 := ctx.URLParam("md5")
	//view := ajaxview.SinglePhotoView(md5)
	view := rootview.SinglePhotoView(md5)
	return view
}

//http://192.168.0.199:8080/delete/  Post method
func (c *RootController) PostDelete(ctx iris.Context) {
	// Accept a path string split by ","
	rawBodyAsBytes, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		fmt.Println("Can not get Delete string.")
	}
	ctx.StatusCode(iris.StatusOK)
	filePathList := strings.Split(string(rawBodyAsBytes), ",")

	model.DeletePhotos(filePathList)
}

//http://192.168.0.199:8080/addtags
func (c *RootController) PostAddtags(ctx iris.Context) {
	// "|" between tags and filePathList
	// "," between filepath in filePathList
	rawBodyAsBytes, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		fmt.Println("Can not get Delete string.")
	}
	ctx.StatusCode(iris.StatusOK)
	tagsWithFiles := string(rawBodyAsBytes) // "tags|filePathList"
	model.AddTags(tagsWithFiles)
}

//Get http://192.168.0.199:8080/slide
func (c *RootController) GetSlide(ctx iris.Context) mvc.Result {
	view := rootview.SlideView()
	return view
}

//Post: http://192.168.0.199:8080/slide for given {"path": /bala/bala/foo.jpg"}
func (c *RootController) PostSlide(ctx iris.Context) {
	type slideFile struct {
		Path string `json:"path"`
	}
	var photo slideFile
	if err := ctx.ReadJSON(&photo); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		//ctx.JSON(iris.Map{"status": iris.StatusBadRequest, "message": err.Error()})
		return
	}
	photoBase64 := model.GenOriginalPicBase64(photo.Path)
	ctx.JSON(iris.Map{"status": iris.StatusOK, "message": photoBase64})
}

//Post: http://192.168.0.199:8080/slideany
func (c *RootController) PostSlideany(ctx iris.Context) {
	getPhoto := model.NextPhoto()
	ctx.JSON(iris.Map{"status": iris.StatusOK, "message": getPhoto()})
}
