package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var path, tpe string
	var err error
	flag.StringVar(&path, "path", "c:/", "file path")
	flag.StringVar(&tpe, "type", "txt", "file type")
	flag.Parse()
	err := traveDir(path, tpe)
	if err != nil {
		fmt.Errorf("open file failed,", err)
	}
}

func traveDir(path string, tpe string) error {
	stat, err := os.Stat(path)
	if stat.IsDir() {
		return nil
	}
	if strings.HasSuffix(stat.Name(), tpe) {
		countLine(path)
	}
	return nil
}

func countLine(path string) (int, error) {
	return 0, nil
}
