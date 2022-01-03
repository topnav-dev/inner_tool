//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || windows
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris windows

package style_format

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

var astyleCmdPath = "./style_format/AStyle"
var astyleArgPath = "./style_format/.astylerc"

func Do(path string) (res string, err error) {
	argString := shellArg()
	if isEmpty(argString...) {
		fmt.Println("Error: astyleArgPath=", astyleArgPath)
	}
	return shell(argString, path)
}

// 执行shell命令
func shell(argString []string, path string) (res string, err error) {
	var execCmd *exec.Cmd
	argString = append(argString, path)
	execCmd = exec.Command(astyleCmdPath, argString...)
	var (
		stdout bytes.Buffer
		stderr bytes.Buffer
	)

	execCmd.Stdout = &stdout
	execCmd.Stderr = &stderr
	err = execCmd.Run()
	res = fmt.Sprintf("Output:\n%s\nError:\n%s", stdout.String(), stderr.String())
	fmt.Print(stdout.String())
	return stdout.String(), err
}

func shellArg() []string {
	file, err := os.Open(astyleArgPath)
	if err != nil {
		log.Fatalf("failed to open")

	}
	// The bufio.NewScanner() function is called in which the
	// object os.File passed as its parameter and this returns a
	// object bufio.Scanner which is further used on the
	// bufio.Scanner.Split() method.
	scanner := bufio.NewScanner(file)
	// The bufio.ScanLines is used as an
	// input to the method bufio.Scanner.Split()
	// and then the scanning forwards to each
	// new line using the bufio.Scanner.Scan()
	// method.
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	// The method os.File.Close() is called
	// on the os.File object to close the file
	file.Close()
	// and then a loop iterates through
	// and prints each of the slice values.
	//for _, lines := range text {
	//	fmt.Println(lines)
	//}
	return text
}

func isEmpty(s ...string) bool {
	ss := path.Join(s...)
	return len(strings.TrimSpace(ss)) == 0
}

func isExist(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
