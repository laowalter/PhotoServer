package rootview

import (
	"fmt"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/photoServer/model"
	"github.com/photoServer/view/util"
)

func RootView(currentPage int64, mobile bool) mvc.Result {
	yearList, err := model.GetYearList()
	if err != nil {
		fmt.Println("Can not Get Year List")
	}

	picList, totalPages, err := model.QueryAllPhotos(currentPage)
	if err != nil {
		fmt.Println("Can not Get Thumbnail List")
	}

	pagers := util.Pagers(currentPage, totalPages)

	template := "index.html"
	if mobile {
		template = "mobileIndex.html"
	}

	return &mvc.View{
		Name: template,
		Data: iris.Map{
			"years":      yearList,
			"thumb":      picList,
			"totalpages": totalPages,
			"pagers":     pagers,
		},
	}
}
