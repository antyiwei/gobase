package main

import (
	"fmt"
)

/*
fmt.Println(1 << (2 * 10
iota是golang语言的常量计数器,只能在常量的表达式中使用。
iota在const关键字出现时将被重置为0(const内部的第一行之前)，const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。
使用iota能简化定义，在定义枚举时很有用。
*/

type ByteSize float64

const (
	_           = iota
	KB ByteSize = (1 << (iota * 10))
	MB          // 1 << (10*2)
	GB          // 1 << (10*3)
	TB          // 1 << (10*4)
	PB          // 1 << (10*5)
	EB          // 1 << (10*6)
	ZB          // 1 << (10*7)
	YB          // 1 << (10*8)
)

func main() {
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
}
