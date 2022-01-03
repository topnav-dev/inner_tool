package dos2unix

import (
	"bytes"
	"regexp"
)

var trailingWhitespace = regexp.MustCompile(`(?m:[\t ]+$)`)

func Do( content []byte) []byte {
	// do magic
	original := content
	content = tidy(content)
	if bytes.Compare(content, original) != 0 {
		return content
	} else {
		return nil
	}
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
