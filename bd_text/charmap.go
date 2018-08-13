package main

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
)

func DecodeWindows1250(enc []byte) string {
	dec := charmap.Windows1250.NewDecoder()
	out, _ := dec.Bytes(enc)
	return string(out)
}

func EncodeWindows1250(inp string) []byte {
	enc := charmap.Windows1250.NewEncoder()
	out, _ := enc.String(inp)
	return []byte(out)
}
func Decode(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, traditionalchinese.Big5.NewDecoder())
	d, e := ioutil.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}
func main() {
	// ibuf := []byte{170, 76, 80, 67}

	// log.Println(DecodeWindows1250(ibuf))

	// a := EncodeWindows1250("你好")
	// b := DecodeWindows1250(a)
	// log.Println(a)
	// log.Println(b)

	// aa := fromWindows1252("你好")
	// log.Println(aa)

	s := []byte{0xB0, 0xAA}
	b, err := Decode(s)
	fmt.Println(string(b))
	fmt.Println(err)

	src := "易伟"
	data, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(src)), simplifiedchinese.GBK.NewEncoder()))

	fmt.Println(data)         //byte
	fmt.Println(string(data)) //打印为乱码

	data2, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader(data), simplifiedchinese.GB18030.NewDecoder()))
	fmt.Println(data2)         //byte
	fmt.Println(string(data2)) //打印为乱码
}

func fromWindows1252(str string) string {
	var arr = []byte(str)
	var buf bytes.Buffer
	var r rune

	for _, b := range arr {
		switch b {
		case 0x80:
			r = 0x20AC
		case 0x82:
			r = 0x201A
		case 0x83:
			r = 0x0192
		case 0x84:
			r = 0x201E
		case 0x85:
			r = 0x2026
		case 0x86:
			r = 0x2020
		case 0x87:
			r = 0x2021
		case 0x88:
			r = 0x02C6
		case 0x89:
			r = 0x2030
		case 0x8A:
			r = 0x0160
		case 0x8B:
			r = 0x2039
		case 0x8C:
			r = 0x0152
		case 0x8E:
			r = 0x017D
		case 0x91:
			r = 0x2018
		case 0x92:
			r = 0x2019
		case 0x93:
			r = 0x201C
		case 0x94:
			r = 0x201D
		case 0x95:
			r = 0x2022
		case 0x96:
			r = 0x2013
		case 0x97:
			r = 0x2014
		case 0x98:
			r = 0x02DC
		case 0x99:
			r = 0x2122
		case 0x9A:
			r = 0x0161
		case 0x9B:
			r = 0x203A
		case 0x9C:
			r = 0x0153
		case 0x9E:
			r = 0x017E
		case 0x9F:
			r = 0x0178
		default:
			r = rune(b)
		}

		buf.WriteRune(r)
	}

	return string(buf.Bytes())
}
