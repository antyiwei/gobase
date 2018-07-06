package main

import "fmt"

// SwitchTest1 熟悉switch
// 可以使用任何类型或表达式作为条件语句
// 不需要写break，一旦条件符合自动终止
func SwitchTest1() {
	a := 1
	switch a {
	case 0:
		fmt.Println("a=0")
	case 1:
		fmt.Println("a=1")

	}
	fmt.Println(a)
}

// SwitchTest2 支持一个初始化表达式（可以是并行方式），右侧需跟分号 左大括号必须和条件语句在同一行
func SwitchTest2() {
	a := 1
	switch {
	case a >= 0:
		fmt.Println("a=0")
	case a >= 1:
		fmt.Println("a=1")

	}
	fmt.Println(a)
}

// SwitchTest3 如希望继续执行下一个case，需使用fallthrough语句
func SwitchTest3() {
	switch a := 1; {

	case a >= 0:
		fmt.Println("a = 0")
	case a >= 1:
		fmt.Println("a = 1")
	}
}
