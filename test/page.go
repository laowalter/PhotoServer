package main

import "fmt"

func main() {
	totalPages := int64(2)
	currentPage := int64(4)
	pagerList := pagers(currentPage, totalPages)
	fmt.Println(pagerList.Before)
	fmt.Println(pagerList.Current)
	fmt.Println(pagerList.After)
}

type PagerList struct {
	Before  []int64
	Current int64
	After   []int64
}

func pagers(currentPage int64, totalPages int64) PagerList {
	var pagers PagerList

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
			for i := int64(1); i <= currentPage; i++ {
				pagers.Before = append(pagers.Before, i)
			}
			for i := currentPage + int64(1); i < currentPage+6 && i <= totalPages; i++ {
				pagers.After = append(pagers.After, i+5)
			}

		} else {
			for i := currentPage - 5; i < currentPage; i++ {
				pagers.Before = append(pagers.Before, i)
			}
			for i := currentPage + 1; i < currentPage+6 && i <= totalPages; i++ {
				pagers.After = append(pagers.After, i+6)
			}

		}
	}
	pagers.Current = currentPage
	return pagers
}
