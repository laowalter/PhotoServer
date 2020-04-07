package ajaxview

import (
	"fmt"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/photoServer/model"
	"github.com/photoServer/view/util"
)

func YearView(thisyear int, currentPage int64, mobile bool) mvc.Result {
	yearList, err := model.GetYearList()
	if err != nil {
		fmt.Println("Can not Get Year List")
	}

	yearPic, totalPages, err := model.QueryPhotosByYear(thisyear, currentPage)
	if err != nil {
		fmt.Println("Can not finding all thumbnail by year")
	}

	pagers := util.Pagers(currentPage, totalPages)

	template := "ajax.html"
	if mobile {
		template = "mobileIndex.html"
	}

	return &mvc.View{
		Name: template,
		Data: iris.Map{
			"years":       yearList,
			"thumb":       yearPic,
			"totalpages":  totalPages,
			"currentyear": thisyear,
			"pagers":      pagers,
		},
	}

}
