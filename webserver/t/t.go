package main

import "fmt"

func main() {
	a := "/home/walter/go/src/github.com/photoServer/testData/DSC00271.JPG"
	b := len("/home/walter")

	fmt.Println(a[12:], b)
}
