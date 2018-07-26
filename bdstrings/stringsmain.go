package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("hello world")

	s := "helloworld"
	substr := "hello"
	fmt.Println(strings.Contains(s, substr))
	fmt.Println("=================")
	{
		/*
		   // 子串substr在s中，返回true
		   func Contains(s, substr string) bool
		   // chars中任何一个Unicode代码点在s中，返回true
		   func ContainsAny(s, chars string) bool
		   // Unicode代码点r在s中，返回true
		   func ContainsRune(s string, r rune) bool
		*/
		fmt.Println(strings.ContainsAny("team", "i"))
		fmt.Println(strings.ContainsAny("failure", "u & i"))
		fmt.Println(strings.ContainsAny("in failure", "s g"))
		fmt.Println(strings.ContainsAny("foo", ""))
		fmt.Println(strings.ContainsAny("", ""))
	}

}
