package version

import (
	"fmt"
	"os"
)

var versionString string
var commitString string

func Do() {
	fmt.Fprintf(os.Stdout, "versionString: %s\r\n", versionString)
	fmt.Fprintf(os.Stdout, "commitString: %s\r\n", commitString)
}
