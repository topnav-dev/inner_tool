//https://github.com/xrstf/txtidy
package main

import (
	"conv/crud"
	"conv/dos2unix"
	"conv/style_format"
	"conv/toUTF8"
	"conv/version"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	flag "github.com/spf13/pflag"
)

// https://learnku.com/articles/53004
// https://github.com/iand/gocookbook/blob/master/recipes/encoding-conversion.md
// https://freshman.tech/snippets/go/check-if-slice-contains-element/
// https://www.gushiciku.cn/pl/pLtt/zh-tw
// https://github.com/gogs/chardet/blob/master/detector_test.go
// https://www.cnblogs.com/sparkdev/p/10833186.html

var (
//dirFlag     = flag.String("dir", "", "root directory to search files in")
//verboseFlag = flag.Bool("v", false, "whether or not to print all visited files")
//allFlag     = flag.Bool("a", false, "run on all files, i.e. do not exclude .git, .hg, node_modules etc.")

//excludedDirs = []string{".git", ".hg", ".svn", "node_modules", "bower_components"}
)

// 定义命令行参数对应的变量
// conv -t dos2unix -v true .c .h
// conv -t all -v true .c .h
var typeFlag = flag.StringP("type", "t", "dos2unix", "toUTF8, dos2unix, all")
var typeExclude = flag.StringP("exclude", "e", "CMSIS", "exclude file")
var verboseFlag = flag.StringP("verbose", "v", "false", "verbose message")

// 定义命令行参数对应的变量
// var cliName = flag.StringP("name", "n", "nick", "Input Your Name")
// var cliAge = flag.IntP("age", "a", 22, "Input Your Age")
// var cliGender = flag.StringP("gender", "g", "male", "Input Your Gender")
// var cliOK = flag.BoolP("ok", "o", false, "Input Are You OK")
// var cliDes = flag.StringP("des-detail", "d", "", "Input Description")
// var cliOldFlag = flag.StringP("badflag", "b", "just for test", "Input badflag")

func wordSepNormalizeFunc(f *flag.FlagSet, name string) flag.NormalizedName {
	from := []string{"-", "_"}
	to := "."
	for _, sep := range from {
		name = strings.Replace(name, sep, to, -1)
	}
	return flag.NormalizedName(name)
}

func flagInit() {
	// 设置标准化参数名称的函数
	flag.CommandLine.SetNormalizeFunc(wordSepNormalizeFunc)
	// 为 age 参数设置 NoOptDefVal
	// flag.Lookup("age").NoOptDefVal = "25"
	// 把 badflag 参数标记为即将废弃的，请用户使用 des-detail 参数
	flag.CommandLine.MarkDeprecated("badflag", "please use --des-detail instead")
	// 把 badflag 参数的 shorthand 标记为即将废弃的，请用户使用 des-detail 的 shorthand 参数
	flag.CommandLine.MarkShorthandDeprecated("badflag", "please use -d instead")
	// 在帮助文档中隐藏参数 gender
	flag.CommandLine.MarkHidden("badflag")
	// 把用户传递的命令行参数解析为对应变量的值
	flag.Parse()
	// fmt.Println("name=", *cliName)
	// fmt.Println("age=", *cliAge)
	// fmt.Println("gender=", *cliGender)
	// fmt.Println("ok=", *cliOK)
	// fmt.Println("des=", *cliDes)
	// fmt.Println("des=", *cliDes)
	fmt.Println("verboseFlag=", *verboseFlag)
	fmt.Println("type=", *typeFlag)
	fmt.Println("Exclude=", *typeExclude)
}

func main() {
	fmt.Println("platform:" + runtime.GOOS + "+" + runtime.GOARCH)
	version.Do()
	flagInit()
	patterns := flag.Args()
	fmt.Println("Tail:", patterns)
	if flag.NArg() == 0 {
		fmt.Println("Error: no file pattern have been given.")
		flag.PrintDefaults()
		os.Exit(1)
	}
	if !(*typeFlag == "toUTF8" || *typeFlag == "dos2unix" || *typeFlag == "all" || *typeFlag == "Astyle") {
		log.Fatal("Error: *typeFlag=", *typeFlag)
		os.Exit(1)
	}

	// check patterns
	for _, pattern := range patterns {
		_, err := filepath.Match(pattern, "dummy")
		if err != nil {
			log.Fatalf("Error: file pattern '%s' is invalid.\n", pattern)
			os.Exit(1)
		}
	}
	files := crud.Search(patterns, *typeExclude)
	for _, file := range files {
		if *typeFlag == "dos2unix" {
			if *verboseFlag == "true" {
			}
			content := crud.Read(file)
			crud.Write(file, dos2unix.Do(content))
		} else if *typeFlag == "toUTF8" {
			content := crud.Read(file)
			crud.Write(file, toUTF8.Do(content))
		} else if *typeFlag == "Astyle" {
			//content := crud.Read(file)
			style_format.Do(file)
		} else if *typeFlag == "all" {
			content := crud.Read(file)
			crud.Write(file, dos2unix.Do(content))
			content = crud.Read(file)
			crud.Write(file, toUTF8.Do(content))
		}
	}
}

//func main() {
//	fmt.Println(runtime.GOOS, "+", runtime.GOARCH)
//	version.Do()
//	flagInit()
//	patterns := flag.Args()
//	fmt.Println("Tail:", patterns)
//}
