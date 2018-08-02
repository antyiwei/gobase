package main

import "fmt"

func main() {
	factorialMain() // 阶乘
	fibonacciMain() // 斐波那契数列
}

// 阶乘
func factorialMain() {

	var i int = 15
	fmt.Printf("%d 的阶乘是 %d\n", i, factorial(uint64(i)))
}

func factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * factorial(n-1)
		return result
	}
	return 1
}

//
func fibonacciMain() {
	var i int
	for i = 0; i < 100; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}
