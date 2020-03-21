package global

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	//Global Defination
)

func ListAllFiles(dirPath string) []string {
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

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func ListFiles(dirPath string, ftype func() []string) ([]string, error) {
	// usage: files, err := global.ListFiles(absPath, global.PICRAW)
	var fileList []string
	err := filepath.Walk(dirPath,
		func(path string, f os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if contains(ftype(), filepath.Ext(strings.ToLower(path))) {
				fileList = append(fileList, path)
			}

			return nil
		})
	if err != nil {
		log.Println(err)
		var emptyArray []string
		return emptyArray, err
	}
	return fileList, nil
}
