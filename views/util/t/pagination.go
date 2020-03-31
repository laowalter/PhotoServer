package util

import (
	"fmt"

	"github.com/photoServer/global"
)

func Pager(pageNo int64) {
	var pageList []int64
	var start int64 = pageNo/(global.PhotosPerPage-int64(1))*(global.PhotosPerPage-int64(1)) + int64(1)
	if pageNo%(global.PhotosPerPage-int64(1)) == 0 {
		start -= global.PhotosPerPage - int64(1)
	}

	for i := start; i < start+global.PhotosPerPage; i++ {
		if i == pageNo {
			fmt.Printf("(%d) ", i)
		} else {
			pageList = append(pageList)
			fmt.Printf("%d ", i)
		}
	}
	fmt.Print("\n")
}
