package rootview

import (
	"fmt"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/photoServer/model"
)

func TagManagement() mvc.Result {

	tags, err := model.ListTags()
	if err != nil {
		fmt.Println("Can not get all the tags form model.ListTags")
	}

	template := "tagmanagement.html"
	return &mvc.View{
		Name: template,
		Data: iris.Map{
			"tags": tags,
		},
	}

}
