//https://github.com/xrstf/txtidy
package main

import (
	"bytes"
	"fmt"
	"github.com/gogs/chardet"
	flag "github.com/spf13/pflag"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"
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
var typeFlag = flag.StringP("type", "t", "dos2unix", "Execute type")
var verboseFlag = flag.BoolP("verbose", "v", true, "verbose message")
var root = getParentDirectory(pwd())
var s = []string{"dos2unix", "dos2unix", "utf8", "Mike"}

func getParentDirectory(dirctory string) string {
	str := func(s string, pos, length int) string {
		runes := []rune(s)
		l := pos + length
		if l > len(runes) {
			l = len(runes)
		}
		return string(runes[pos:l])
	}(dirctory, 0, strings.LastIndex(dirctory, "/"))
	return str
}

func wordSepNormalizeFunc(f *flag.FlagSet, name string) flag.NormalizedName {
	from := []string{"-", "_"}
	to := "."
	for _, sep := range from {
		name = strings.Replace(name, sep, to, -1)
	}
	return flag.NormalizedName(name)
}

func flag_init() {
	// 设置标准化参数名称的函数
	flag.CommandLine.SetNormalizeFunc(wordSepNormalizeFunc)
	// 为 age 参数设置 NoOptDefVal
	flag.Lookup("age").NoOptDefVal = "25"
	// 把 badflag 参数标记为即将废弃的，请用户使用 des-detail 参数
	flag.CommandLine.MarkDeprecated("badflag", "please use --des-detail instead")
	// 把 badflag 参数的 shorthand 标记为即将废弃的，请用户使用 des-detail 的 shorthand 参数
	flag.CommandLine.MarkShorthandDeprecated("badflag", "please use -d instead")
	// 在帮助文档中隐藏参数 gender
	flag.CommandLine.MarkHidden("badflag")
	// 把用户传递的命令行参数解析为对应变量的值
	flag.Parse()
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func pwd() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func Native_conv() string {
	cmd, err := exec.Command("/bin/sh", "/path/to/file.sh").Output()
	if err != nil {
		fmt.Printf("error %s", err)
	}
	output := string(cmd)
	return output
}

func main() {
	version()
	flag.Parse()
	patterns := flag.Args()
	fmt.Printf("Tail: %+q\n", patterns)
	fmt.Printf("verboseFlag=%t\n", *verboseFlag)
	fmt.Println("name=", *typeFlag)
	time.Sleep(5 * time.Second)

	if len(patterns) == 0 {
		fmt.Printf("Error: no file pattern have been given.\n\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// check patterns
	for _, pattern := range patterns {
		_, err := filepath.Match(pattern, "dummy")
		if err != nil {
			fmt.Printf("Error: file pattern '%s' is invalid.\n", pattern)
			os.Exit(1)
		}
	}

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
	for _, file := range files {
		if *verboseFlag {
			fmt.Printf("%s\n", file)
		}
		if *typeFlag == "dos2unix" {
			dos2unix(file)
		}
		if *typeFlag == "all" {
			buffer := make([]byte, 32<<10)
			f, err := os.Open(filepath.Join(file))
			if err != nil {
				fmt.Println(err)
			}
			size, _ := io.ReadFull(f, buffer)
			input := buffer[:size]

			textDetector := chardet.NewTextDetector()
			//result, err := textDetector.DetectBest(input)
			//if err != nil {
			//	fmt.Println(err)
			//}
			//if result.Charset != "UTF-8" {
			//	fmt.Printf("Expected charset %s, actual %s\n", "UTF-8", result.Charset)
			//}
			result, err := textDetector.DetectAll(input)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(result)
			}

		}
	}
}

var trailingWhitespace = regexp.MustCompile(`(?m:[\t ]+$)`)

func dos2unix(path string) error {
	// read file into memory
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf(" error: %s\n", err.Error())
		return nil
	}

	// do magic
	original := content
	content = tidy(content)

	if bytes.Compare(content, original) != 0 {
		stats, _ := os.Stat(path)
		err := ioutil.WriteFile(path, content, stats.Mode())
		if err != nil {
			fmt.Printf(" error: %s\n", err.Error())
		} else {
			fmt.Printf(" fixed.\n")
		}
	}
	return nil
}

func tidy(content []byte) []byte {
	// turn Windows newlines into Unix newlines
	content = bytes.Replace(content, []byte{'\r'}, []byte{}, -1)

	// remove UTF BOMs
	if len(content) >= 3 && bytes.Equal(content[0:3], []byte("\xEF\xBB\xBF")) {
		content = content[3:]
	}

	// trim trailing whitespace in each line
	content = trailingWhitespace.ReplaceAllLiteral(content, []byte{})

	// trim leading and trailing file space
	content = bytes.TrimSpace(content)

	// and make sure the file ends with a newline character
	content = append(content, '\n')

	return content
}
