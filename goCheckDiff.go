package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"time"
)

var ignoreList = []string{".git", ".idea", ".DS_Store"}

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Usage: " + filepath.Base(os.Args[0]) + " [base_directory_path] [target_directory_path]")
		return
	}

	baseDir := addSlashAtEnd(os.Args[1])
	targetDir := addSlashAtEnd(os.Args[2])

	fmt.Printf("%s Start. \n", time.Now().Format("2006-01-02 15:04:05"))

	isDifferent := false
	err := checkDiff(baseDir, targetDir, &isDifferent)
	if err != nil {
		fmt.Println(err)
		return
	}

	if !isDifferent {
		fmt.Println("All files are same.")
	}

	fmt.Printf("%s End. \n", time.Now().Format("2006-01-02 15:04:05"))
}

func checkDiff(baseDir, targetDir string, isDifferent *bool) error {

	files, err := ioutil.ReadDir(baseDir)
	if err != nil {
		return err
	}

	for _, file := range files {

		if arrayContains(ignoreList, file.Name()) {
			continue
		}

		if file.IsDir() {
			// Check existence of Directory
			checkbaseDir := baseDir + file.Name() + "/"
			checktargetDir := targetDir + file.Name() + "/"

			if _, err := os.Stat(checktargetDir); err != nil {
				fmt.Printf("Directory not found. -> %s \n", checktargetDir)
				*isDifferent = true
				continue
			}

			// Recursion
			checkDiff(checkbaseDir, checktargetDir, isDifferent)
		}

		checkFromFile := baseDir + file.Name()
		checkToFile := targetDir + file.Name()

		// Check existence of File
		if _, err := os.Stat(checkToFile); err != nil {
			fmt.Printf("File not found. -> %s \n", checkToFile)
			*isDifferent = true
			continue
		}
		// Check Files difference
		fileFrom, err := ioutil.ReadFile(checkFromFile)
		if err != nil {
			continue
		}
		fileTo, err := ioutil.ReadFile(checkToFile)
		if err != nil {
			continue
		}
		if !reflect.DeepEqual(fileFrom, fileTo) {
			fmt.Printf("File is different. -> %s \n", checkToFile)
			*isDifferent = true
			continue
		}
	}

	return nil
}

func arrayContains(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

func addSlashAtEnd(str string) string {
	if str[len(str)-1:len(str)] != "/" {
		str += "/"
	}
	return str
}
