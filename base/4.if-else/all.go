package main

import "fmt"

/*
跳转语句goto, break, continue

三个语法都可以配合标签使用
标签名区分大小写，若不使用会造成编译错误
Break与continue配合标签可用于多层循环的跳出
Goto是调整执行位置，与其它2个语句配合标签的结果并不相同

*/

// GotoTest1 利用goto跳出循环
func GotoTest1() {
LABEL:
	for i := 0; i < 10; i++ {
		if i > 2 {
			break LABEL
		} else {
			fmt.Println(i)
		}
	}

}

// GotoTest2 利用goto跳出多重循环
func GotoTest2() {
LABEL:
	for i := 0; i < 10; i++ {
		for {
			fmt.Println(i)
			continue LABEL
		}
	}
}
