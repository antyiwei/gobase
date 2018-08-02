package main

import "fmt"

func main() {

	str := "abcdefghijk"
	strsub := "a"

	num := NaiveStringMatcher(str, strsub)
	if num >= 0 {
		fmt.Println("Pattern occurs with shift:", num)
	} else {
		fmt.Println("Not macth the P in the T.")
	}

}

func NaiveStringMatcher(str string, strsub string) int {
	n := len(str)
	m := len(strsub)

	strRune := []rune(str)
	strsubRune := []rune(strsub)
	if n < m {
		return -1
	}
	var s_length int = n - m

	for i := 0; i <= s_length; i++ {
		var flag bool = false
		for j := 0; j < m; j++ {
			if strRune[i+j] == strsubRune[j] {
				flag = true
			} else {
				flag = false
			}
		}

		if flag {
			return i
		}
	}
	return -1

}
