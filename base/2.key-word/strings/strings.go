package main

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
)

func stringsToJSON(str string) string {
	var jsons bytes.Buffer
	for _, r := range str {
		rint := int(r)
		if rint < 128 {
			jsons.WriteRune(r)
		} else {
			jsons.WriteString("\\u")
			jsons.WriteString(strconv.FormatInt(int64(rint), 16))
		}
	}
	return jsons.String()
}
func StringsToJSON2(str string) string {
	rs := []rune(str)
	jsons := ""
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			jsons += string(r)
		} else {
			jsons += "\\u" + strconv.FormatInt(int64(rint), 16) // json
		}
	}
	return jsons
}

func main() {
	fmt.Println("=============", time.Now().Format("2006-01-02 15:04:05.999"))
	name := "hello world!"
	for index := 0; index < 1000; index++ {
		name += StringsToJSON2(name)
	}
	fmt.Println("=============", time.Now().Format("2006-01-02 15:04:05.999"))
}
