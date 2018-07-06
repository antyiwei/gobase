package main

import "fmt"

func main() {
	// s := []int{1, 2, 3, 4}
	// C(s)
	// fmt.Println(s)
	//=======================
	// a := 2
	// C1(a)
	// fmt.Println(a)
	//=======================
	// b := C2
	// b()
	//=======================
	// c := func() {
	// 	fmt.Println("Func 匿名函数")
	// }
	// c()
	//=======================
	// d := closure(10)
	// fmt.Println(d(1))
	// fmt.Println(d(2))
	//=======================

	// fmt.Println("a")
	// defer fmt.Println("b")
	// defer fmt.Println("c")

	// DeferA()

	// DeferB()

	DeferC()
	DeferC1()
	DeferC2()
}

// closure 封闭包函数
func closure(x int) func(int) int {
	fmt.Printf("%p\n", &x)
	return func(y int) int {
		fmt.Printf("%p\n", &x)
		return x + y
	}
}

// Work 作业
func Work() {
	var fs = [4]func(){}
	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i = ", i)
		defer func() {
			fmt.Println("defer_closure i = ", i)
		}()
		fs[i] = func() {
			fmt.Println("closure i = ", i)
		}
	}
	for _, f := range fs {
		f()
	}
}

// A 函数
func A() (a, b, c int) {
	a, b, c = 1, 2, 3
	return
}

// B 不定长变参
func B(a ...int) {
	fmt.Println(a)
}

// C 是否是值类型还是引用类型
func C(s []int) {
	s[0] = 5
	s[1] = 6
	s[2] = 7
	s[3] = 8
	fmt.Println(s)
}

// C1 值类型
func C1(a int) {
	a = 9
	fmt.Println(a)
}

// C2 函数作为类型
func C2() {
	fmt.Println("Func C2")
}

// DeferA 先进后出，后进先出
func DeferA() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println("----------------------------")
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	// fmt.Println("----------------------------")
	// for i := 0; i < 5; i++ {
	// 	defer func() {
	// 		fmt.Println(i)
	// 	}()
	// }
}

// DeferB  defer 在闭包的时候请多注意
func DeferB() {
	// fmt.Println("----------------------------")
	for i := 0; i < 5; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}

// DeferC 函数
func DeferC() {
	fmt.Println("func DeferC")
}

// DeferC1 函数
func DeferC1() {

	/*
		第二种情况output:
				func DeferC
				Recover in B
				func DeferC2
	*/
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover in B")
		}
	}()
	/*
		第一种情况output:
		func DeferC
		panic: Panic in DeferC1

		goroutine 1 [running]:
		main.DeferC1()
		        /Users/antyiwei/mygo/src/github.com/antyiwei/gobase/base/6.函数/func.go:130 +0x39
		main.main()
		        /Users/antyiwei/mygo/src/github.com/antyiwei/gobase/base/6.函数/func.go:36 +0x25
		exit status 2
	*/
	panic("Panic in DeferC1")

}

// DeferC2 函数 DeferC2
func DeferC2() {
	fmt.Println("func DeferC2")
}
