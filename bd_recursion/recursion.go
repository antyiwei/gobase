package main

import (
	"fmt"
	"math/big"
	"time"

	"github.com/antyiwei/gobase/utils/matrix"
)

/* 斐波那契  */
const MaxNum = 1000000

var fibs [MaxNum]*big.Int
var facs [MaxNum]*big.Int

func main() {
	factorialMain() // 阶乘
	//fibonacciMain() // 斐波那契数列

	//FibMain() // 使用矩阵求值
}

// 阶乘
func factorialMain() {
	var i int
	for i = 0; i < MaxNum; i++ {
		fmt.Printf("factorial(%d) 的阶乘是:%d\n", i, factorial(i))
	}
}

func factorial(n int) (res *big.Int) {
	if n < 2 {
		res = big.NewInt(1)
	} else {
		temp := new(big.Int)
		nint64 := big.NewInt(int64(n))
		res = temp.Mul(nint64, facs[n-1]) //n * Factorial(n-1)
	}
	facs[n] = res
	return
}

// 斐波那契
func fibonacciMain() {
	start := time.Now()
	var i int
	for i = 0; i < MaxNum; i++ {

		fmt.Printf("fibonacci(%d)is:%d\n", i, fibonacci(i))
	}
	end := time.Now()
	delete := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time:%s\n", delete)
}

func fibonacci(n int) (res *big.Int) {
	// if fibs[n].Sign() != 0 {
	// 	res = fibs[n]
	// 	return
	// }
	if n < 2 {
		res = big.NewInt(1)
	} else {
		temp := new(big.Int)
		res = temp.Add(fibs[n-1], fibs[n-2])
	}
	fibs[n] = res
	return
}

// 斐波那契 2种方法
//var fibarry = [3]int{0, 1, 0}
//
//func fibonacci2(n int) int {
//	for i := 2; i <= n; i++ {
//		fibarry[2] = fibarry[0] + fibarry[1]
//		fibarry[0] = fibarry[1]
//		fibarry[1] = fibarry[2]
//	}
//	return fibarry[2]
//}

// 三种，矩证
func FibMain() {
	start := time.Now()
	n := MaxNum - 1
	m := Fib(n)
	fmt.Printf("f(%d)的结果是:%d", n, m)
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)

}

// 求矩阵的n次幂
func MatPow(a matrix.Matrix, b int) matrix.Matrix {
	arr0 := [4]*big.Int{big.NewInt(1), big.NewInt(0), big.NewInt(0), big.NewInt(1)}
	s := matrix.New(2, 2, arr0[0:4])

	for b > 0 {
		if b&1 == 1 {
			s = *matrix.Multiply(s, a)
			b = b >> 1
		} else {
			b = b >> 1
		}
		a = *matrix.Multiply(a, a)
	}
	return s
}

func Fib(n int) *big.Int {
	arr0 := [6]*big.Int{big.NewInt(1), big.NewInt(1), big.NewInt(1), big.NewInt(0), big.NewInt(2), big.NewInt(1)}
	k := matrix.New(2, 2, arr0[0:4])
	s := MatPow(k, n)
	d := matrix.New(2, 1, arr0[0:2])
	s = *matrix.Multiply(s, d)
	return s.Get(2, 1)

}
