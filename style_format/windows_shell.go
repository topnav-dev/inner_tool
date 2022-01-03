//go:build windows
// +build windows
package style_format

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

func Astyle(){
	shell
}

// 执行shell命令
func shell(cmd string) (res string, err error) {
	var execCmd *exec.Cmd
	if runtime.GOOS == "windows" {
		execCmd = exec.Command("cmd.exe", "/c", cmd)
	} else {
		execCmd = exec.Command("bash", "-c", cmd)
	}
	var (
		stdout bytes.Buffer
		stderr bytes.Buffer
	)

	execCmd.Stdout = &stdout
	execCmd.Stderr = &stderr
	err = execCmd.Run()
	res = fmt.Sprintf("Output:\n%s\nError:\n%s", stdout.String(), stderr.String())
	return res, err
}

func shell_arg(){
	file, err := os.Open(".astylerc")
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
		text = append(text, scanner.Text() + " ")
	}
	// The method os.File.Close() is called
	// on the os.File object to close the file
	file.Close()
	// and then a loop iterates through
	// and prints each of the slice values.
	//for _, lines := range text {
	fmt.Println(text)
	//}
}
