package main

import (
	"log"
	"os"
	"path/filepath"
)

func ListAllFilesTree(dirPath string) []string {
	var fileList []string
	err := filepath.Walk(dirPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			//fmt.Println(path, info.Size())
			fileList = append(fileList, path)
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	return fileList
}

func ListPicVideoFilesTree(dirPath string) []string {
	ext := ".jpg"
	var fileList []string
	err := filepath.Walk(dirPath,
		func(path string, f os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if filepath.Ext(path) == ext {
				fileList = append(fileList, path)
			}

			return nil
		})
	if err != nil {
		log.Println(err)
	}
	return fileList
}

/*
func main() {
	path := "/home/walter/go/src/github.com/photoServer/testData/"
	files := ListAllFilesTree(path)
	for _, file := range files {
		fmt.Println(file)
	}

}
*/
