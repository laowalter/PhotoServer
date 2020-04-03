package util

import (
	"github.com/photoServer/global"
)

func Pagers(currentPage int64, totalPages int64) global.Pagers {
	var pagers global.Pagers

	if currentPage > totalPages {
		return pagers
	}

	if totalPages < int64(11) {
		for i := int64(1); i < currentPage; i++ {
			pagers.Before = append(pagers.Before, i)
		}
		for i := currentPage + 1; i <= totalPages; i++ {
			pagers.After = append(pagers.After, i)
		}
	} else {
		if currentPage < int64(6) {
			for i := int64(1); i < currentPage; i++ {
				pagers.Before = append(pagers.Before, i)
			}
			for i := currentPage + int64(1); i < currentPage+6 && i <= totalPages; i++ {
				pagers.After = append(pagers.After, i)
			}

		} else {
			for i := currentPage - 5; i < currentPage; i++ {
				pagers.Before = append(pagers.Before, i)
			}
			for i := currentPage + 1; i < currentPage+6 && i <= totalPages; i++ {
				pagers.After = append(pagers.After, i)
			}

		}
	}

	if currentPage == int64(1) {
		pagers.Before = []int64{}
	}
	if currentPage == totalPages {
		pagers.After = []int64{}
	}
	pagers.Current = currentPage
	return pagers
}
