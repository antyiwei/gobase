package main

import "fmt"

/*

> Output:
command-line-arguments
A
> Elapsed: 0.306s
> Result: Success
*/
func main() {
	var a int = 65
	b := string(a)
	fmt.Println(b)
}
