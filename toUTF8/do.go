package toUTF8

import (
	"fmt"
	"github.com/gogs/chardet"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
	"strings"
)

func Do(content []byte) []byte{
	textDetector := chardet.NewTextDetector()
	result, err := textDetector.DetectBest(content)
	if err != nil {
		fmt.Println(err)
	}
	result.Charset = strings.ToUpper(result.Charset)
	if !(result.Charset == "ISO-8859-1" ||
		result.Charset == "ISO-8859-2" ||
		result.Charset == "ISO-8859-9" ||
		result.Charset == "BIG5" ||
		result.Charset == "GB18030" ||
		result.Charset == "UTF-8") {
		fmt.Printf("Expected charset %s, actual %s\n", "UTF-8", result.Charset)
	}
	if result.Charset == "BIG5" {
		Big5toUTF8 := traditionalchinese.Big5.NewDecoder()
		utf8, _, _ := transform.Bytes(Big5toUTF8, content)
		return utf8
	}
	if result.Charset == "GB18030" {
		GB18030toUTF8 :=simplifiedchinese.GB18030.NewDecoder()
		utf8, _, _ := transform.Bytes(GB18030toUTF8, content)
		return utf8
	}
	return nil
}

