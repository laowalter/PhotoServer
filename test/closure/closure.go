package main

import "fmt"

func f() func() int {
	i := 0
	return func() int {
		i++
		fmt.Println(i)
		return i
	}
}

func main() {
	c1 := f()
	c1() // 打印 1
	c1 = f()
	c1() // 打印 1
	c1 = f()
	c1() // 打印 1
}
