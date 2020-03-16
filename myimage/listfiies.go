package myimage

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/photoServer" //Global Defination
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

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func ListPicVideoFilesTree(dirPath string) []string {
	var fileList []string
	err := filepath.Walk(dirPath,
		func(path string, f os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if contains(photoServer.PICVIDEO(), filepath.Ext(strings.ToLower(path))) {
				fileList = append(fileList, path)
			}

			return nil
		})
	if err != nil {
		log.Println(err)
	}
	return fileList
}

func ListPics(dirPath string) []string {
	var fileList []string
	err := filepath.Walk(dirPath,
		func(path string, f os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if contains(photoServer.PIC(), filepath.Ext(strings.ToLower(path))) {
				fileList = append(fileList, path)
			}

			return nil
		})
	if err != nil {
		log.Println(err)
	}
	return fileList
}
