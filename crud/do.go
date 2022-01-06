package crud

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Search(patterns []string, exclude string) []string {
	// add file to files
	var root = getWorkingDirPath()
	fmt.Println("Search path:", root)
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}
		for _, string := range patterns {
			if !info.IsDir() && filepath.Ext(path) == string {
				if strings.Contains(path, exclude) {
					//fmt.Println("Excluded", path)
				} else {
					files = append(files, path)
				}
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func Read(path string) []byte {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	return b
}

func Write(path string, content []byte) error {
	// empty file
	if content == nil {
		return nil
	}
	err := write(path, content)
	// write error
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
	return err
}

func write(path string, content []byte) error {
	stats, _ := os.Stat(path)
	return ioutil.WriteFile(path, content, stats.Mode())
}

func getParentDirectory(wd string) string {
	parent := filepath.Dir(wd)
	return parent
}

// BaseOn exe
func getExePath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exePath := filepath.Dir(ex)
	return exePath
}

// Base on exe
func getAbsPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	return dir
}

// Base on command line path
func getWorkingDirPath() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return dir
}