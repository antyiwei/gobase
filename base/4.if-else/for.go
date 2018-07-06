package main

import "fmt"

// ForTest1 测试1
func ForTest1() {
	a := 1
	for {
		a++
		if a > 3 {
			break
		}
	}
	fmt.Println(a)
}

// ForTest2 测试2
func ForTest2() {
	a := 1
	for a <= 3 {
		a++
	}
	fmt.Println(a)
}

// ForTest3 初始化和步进表达式可以是多个值
func ForTest3() {
	a := 1
	for i := 0; i < 3; i++ {
		a++
	}
	fmt.Println(a)
}
