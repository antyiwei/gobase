package main

import "fmt"

/*
IfTest1 > Output:
command-line-arguments
小于等于6
1
true
> Elapsed: 0.317s
> Result: Success
*/
func IfTest1() {
	a := true
	if a, b, c := 1, 2, 3; a+b+c > 6 {
		fmt.Println("大于6")
	} else {
		fmt.Println("小于等于6")
		fmt.Println(a)
	}
	fmt.Println(a)
}
