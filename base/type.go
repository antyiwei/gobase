package main

import (
	"fmt"
)

/*

> Output:
command-line-arguments
false 0 0
> Elapsed: 0.310s
> Result: Success
*/
func main() {
	var isShow bool

	var age int

	var num int8

	fmt.Println(isShow, age, num)

	// 多个变量的声明与赋值

	//多个变量的声明
	var a, b, c, d int
	// 多个变量的赋值
	a, b, c, d = 1, 2, 3, 4

	// 多个变量声明的同时赋值
	var e, f, g, h int = 5, 6, 7, 8
	// 省略变量类型，由系统推断
	var i, j, k, l = 9, 10, 11, 12
	// 多个变量声明与赋值的最简写法
	i, m, n, o := 12, 23, 434, 45

	// 在相互兼容的两种类型之间进行转换
	var a float64 = 1.20
	b := int(a)

	// 以下表达式无法通过编译
	var c bool = true
	d := int(c)

}

// 定义单个常量
const a int = 1
const b = 'A'
const (
	text   = "123"
	length = len(test)
	num    = b * 20
)

// 同时定义多个变量
const i, j, k = 1, "2", '3'
const (
	text2, length2, num2 = "234", len(text2), k * 10
)
