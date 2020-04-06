package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	// fileName := "birthday_001.txt"
	// // => Birthday - 1 of 4 .txt
	// newName, err := match(fileName, 4)
	// if err != nil {
	// 	fmt.Println("No mat ch!")
	// 	os.Exit(1)
	// }
	// fmt.Println(newName)

	dir := "./sample"
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	count := 0

	var toRename []string
	for _, file := range files {
		if file.IsDir() {
		} else {
			_, err := match(file.Name(), 4)
			if err == nil {
				count++
				toRename = append(toRename, file.Name())
			}
		}
	}

	for _, origFilename := range toRename {
		origPath := filepath.Join(dir, origFilename)
		newFilename, err := match(origFilename, count)
		if err != nil {
			fmt.Println(toRename)
			panic(err)
			//continue
		}
		newPath := filepath.Join(dir, newFilename)
		err = os.Rename(origPath, newPath)
		if err != nil {
			panic(err)
		}
		fmt.Printf("mv %s => %s\n", origPath, newPath)
	}
}

// match returns the new file name or an error
func match(filename string, total int) (string, error) {
	//split the filename parts
	pieces := strings.Split(filename, ".")
	ext := pieces[len(pieces)-1]
	tmp := strings.Join(pieces[0:len(pieces)-1], ".")
	pieces = strings.Split(tmp, "_")
	name := strings.Join(pieces[0:len(pieces)-1], "_")
	number, err := strconv.Atoi(pieces[len(pieces)-1])
	if err != nil {
		return "", fmt.Errorf("%s didn't match our pattern ", filename)
	}
	return fmt.Sprintf("%s - %d of %d.%s", strings.Title(name), number, total, ext), nil
}
