package main

import (
	"fmt"

	"github.com/photoServer/views/util"
)

func main() {
	fmt.Println("vim-go")
	p := util.Pagers(1, 26)
	fmt.Println(p.Before)
	fmt.Println(p.Current)
	fmt.Println(p.After)
}
