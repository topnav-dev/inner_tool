package crud

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var root = getParentDirectory()

func Search(patterns []string) []string {
	// add file to files
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}
		for _, string := range patterns {
			if !info.IsDir() && filepath.Ext(path) == string {
				files = append(files, path)
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func Read(path string) []byte{
	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf(" error: %s\n", err)
	}
	return b
}

func Write(path string, content []byte) error{
	// empty file
	if content == nil{
		return nil
	}
	err := write(path, content)
	// write error
	if err != nil{
		fmt.Printf("error: %s\n", err.Error())
	}
	return err
}

func write(path string, content []byte) error{
	stats, _ := os.Stat(path)
	return ioutil.WriteFile(path, content, stats.Mode())
}

func getParentDirectory() string {
	wd,err := os.Getwd()
	if err != nil {
		panic(err)
	}
	parent := filepath.Dir(wd)
	return parent
}

func init() {

}
