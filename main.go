package main

import (
	"fmt"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/recover"
	"github.com/photoServer/model"

	"github.com/kataras/iris/mvc"
)

func main() {
	app := iris.New()
	app.Use(recover.New())
	app.Logger().SetLevel("debug")

	app.HandleDir("/static", "./assets")

	tmpl := iris.HTML("./views", ".html")
	app.RegisterView(tmpl)

	mvc.Configure(app.Party("/basic"), basicMVC)

	app.Run(iris.Addr(":8080"))
}

type basicController struct {
}

func basicMVC(app *mvc.Application) {
	app.Router.Use(func(ctx iris.Context) {
		ctx.Application().Logger().Infof("Path: %s", ctx.Path())
		ctx.Next()
	})

	app.Handle(new(basicController))

}

func (c *basicController) Get() mvc.Result {
	YearList, err := model.GetYearList()
	if err != nil {
		fmt.Println("Can not Get Year List")
	}

	PicList, err := model.GetThumbList()
	if err != nil {
		fmt.Println("Can not Get Thumbnail List")
	}

	return mvc.View{
		Name: "index.html",
		Data: iris.Map{"years": YearList, "thumb": PicList},
	}
}

func (c *basicController) GetHosts() string {
	body := fmt.Sprintf("Hello visits from you: %d", 1)
	return body
}
