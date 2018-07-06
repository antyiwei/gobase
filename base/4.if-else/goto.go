package main

import "fmt"

// GotoTest3 死循环，小心执行
func GotoTest3() {
LABEL:
	for i := 0; i < 10; i++ {
		for {
			fmt.Println(i)
			goto LABEL
		}
	}
}
