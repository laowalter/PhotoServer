package rootview

import (
	"fmt"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/photoServer/model"
	"github.com/photoServer/view/util"
)

func YearView(thisyear int, currentPage int64) mvc.Result {
	yearList, err := model.GetYearList()
	if err != nil {
		fmt.Println("Can not Get Year List")
	}

	yearPic, totalPages, err := model.QueryPhotosByYear(thisyear, currentPage)
	if err != nil {
		fmt.Println("Can not finding all thumbnail by year")
	}

	pagers := util.Pagers(currentPage, totalPages)
	return &mvc.View{
		Name: "index.html",
		Data: iris.Map{
			"years":       yearList,
			"thumb":       yearPic,
			"totalpages":  totalPages,
			"currentyear": thisyear,
			"pagers":      pagers,
		},
	}

}
