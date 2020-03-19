package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	paths := []string{
		"/a/b/c",
		"/b/c",
		"./b/c",
	}

	fmt.Println("On Unix:")
	for _, p := range paths {
		abs, err := filepath.Abs(p)
		fmt.Printf("%q: %q %v\n", p, abs, err)
	}

}
