package main

import (
	"fmt"
	"os"
)

var versionString string
var commitString string

func version() {
	fmt.Fprintf(os.Stdout, "versionString: %s\r\n", versionString)
	fmt.Fprintf(os.Stdout, "commitString: %s\r\n", commitString)
}
