package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

// used to find same file named ...(1).txt, ...(2).xls
func main() {
	var path string
	// go run deleteSameFile.go -path ...
	flag.StringVar(&path, "path", "", "please input path to search.")
	if err := findSameNameFiles(path); err != nil {
		fmt.Println("error")
	}
}

func VisitFile(path string, f os.FileInfo, err error) (e error) {
	matched, err := regexp.MatchString("\\([0-9]?\\)?.mp3", path)
	if err != nil {
		fmt.Println("file error", path)
	}
	if matched == true {
		fmt.Println(path)
		err = os.Remove(path)
		if err != nil {
			fmt.Println("delete file error", err)
		}
	}
	return err
}

func findSameNameFiles(path string) error {
	err := filepath.Walk(path, VisitFile)
	return err
}
