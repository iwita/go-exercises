package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// type file struct {
// 	path string
// 	name string
// }

// -----RegExp-----------
// ^ 	means starts with
// ()	define a group
// . 	means match any character
// * 	means match something as many times you want
// ? 	means do the mathing for the smallest number of characters
// + 	means that there gonna be at least one of those

// example to match "family easter 2014 (22 of 25).txt"
// rename it into "2014 - family easter - 22 of 55.txt"

var re = regexp.MustCompile("^(.+?) ([0-9]{4}) [(]([0-9]+) of ([0-9]+)[)][.](.+?)$")
var replaceString = "$2 - $1 - $3 of $4.$5"

func main() {
	// var dry bool
	// flag.BoolVar(&dry, "dry", true, "whether or not this should be a real or a dry run")
	// flag.Parse()
	// walkDir := "sample"
	// // map how many of each name found
	// toRename := make(map[string][]string)
	// filepath.Walk(walkDir, func(path string, info os.FileInfo, err error) error {
	// 	if info.IsDir() {
	// 		return nil
	// 	}
	// 	curDir := filepath.Dir(path)
	// 	if m, err := match(info.Name()); err == nil {
	// 		key := filepath.Join(curDir, fmt.Sprintf("%s.%s", m.base, m.ext))
	// 		toRename[key] = append(toRename[key], info.Name())
	// 	}
	// 	return nil
	// })
	// for key, files := range toRename {
	// 	dir := filepath.Dir(key)
	// 	n := len(files)
	// 	sort.Strings(files)
	// 	for i, filename := range files {
	// 		res, _ := match(filename)
	// 		newFilename := fmt.Sprintf("%s-%d of %d.%s", res.base, (i + 1), n, res.ext)
	// 		oldPath := filepath.Join(dir, filename)
	// 		newPath := filepath.Join(dir, newFilename)
	// 		fmt.Printf("mv %s => %s\n", oldPath, newPath)
	// 		if !dry {
	// 			err := os.Rename(oldPath, newPath)
	// 			if err != nil {
	// 				fmt.Println("Error renaming: ", oldPath, err.Error())
	// 			}
	// 		}
	// 	}
	// }

	var dry bool
	flag.BoolVar(&dry, "dry", true, "whether or not this should be a real or a dry run")
	flag.Parse()
	walkDir := "sample"
	// map how many of each name found
	var toRename []string
	filepath.Walk(walkDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if _, err := match(info.Name()); err != nil {
			toRename = append(toRename, path)
		}
		return nil
	})

	for _, oldPath := range toRename {
		dir := filepath.Dir(oldPath)
		filename := filepath.Base(oldPath)
		newFilename, _ := match_reg(filename)
		newPath := filepath.Join(dir, newFilename)
		fmt.Printf("mv %s => %s\n", oldPath, newPath)
		if !dry {
			err := os.Rename(oldPath, newPath)
			if err != nil {
				fmt.Println("Error renaming: ", oldPath, err.Error())
			}
		}
	}
	// for key, files := range toRename {
	// 	dir := filepath.Dir(key)
	// 	n := len(files)
	// 	sort.Strings(files)
	// 	for i, filename := range files {
	// 		res, _ := match(filename)
	// 		newFilename := fmt.Sprintf("%s-%d of %d.%s", res.base, (i + 1), n, res.ext)
	// 		oldPath := filepath.Join(dir, filename)
	// 		newPath := filepath.Join(dir, newFilename)
	// 		fmt.Printf("mv %s => %s\n", oldPath, newPath)
	// 		if !dry {
	// 			err := os.Rename(oldPath, newPath)
	// 			if err != nil {
	// 				fmt.Println("Error renaming: ", oldPath, err.Error())
	// 			}
	// 		}
	// 	}
	// }
}

type matchResult struct {
	base  string
	index int
	ext   string
}

// match returns the new file name or an error
func match(filename string) (*matchResult, error) {
	//split the filename parts
	pieces := strings.Split(filename, ".")
	ext := pieces[len(pieces)-1]
	tmp := strings.Join(pieces[0:len(pieces)-1], ".")
	pieces = strings.Split(tmp, "_")
	name := strings.Join(pieces[0:len(pieces)-1], "_")
	number, err := strconv.Atoi(pieces[len(pieces)-1])
	if err != nil {
		return nil, fmt.Errorf("%s didn't match our pattern ", filename)
	}
	return &matchResult{
		strings.Title(name),
		number,
		ext,
	}, nil
}

func match_reg(filename string) (string, error) {

	if !re.MatchString(filename) {
		return "", fmt.Errorf("%s didn't match our pattern ", filename)
	}
	return re.ReplaceAllString(filename, replaceString), nil
}
