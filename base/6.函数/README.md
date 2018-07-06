# 函数

## 函数function

- Go 函数 不支持 嵌套、重载和默认参数
- 但支持以下特性：
    - 无需声明原型、不定长度变参、多返回值、命名返回值参数
    - 匿名函数、闭包
- 定义函数使用关键字 func，且左大括号不能另起一行
- 函数也可以作为一种类型使用

```go
func A() (a, b, c int) {
	a, b, c = 1, 2, 3
	return
}
// B 不定长变参
func B(a ...int) {
	fmt.Println(a)
}
```

如果传入的引用类型，在函数中改变值，也修改原有的值;如果是值类型，函数中改变后，不会修改值
```go
package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4}
	C(s)
	fmt.Println(s)

	a := 2
	C1(a)
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

//  C1 值类型
func C1(a int) {
	a = 9
	fmt.Println(a)
}

/*
out:
[5 6 7 8]
[5 6 7 8]
9
2
*/
```

函数也是类型
```go
package main

import "fmt"

func main() {
	b := C2
	b()
}

// C2 函数作为类型
func C2() {
	fmt.Println("Func C2")
}
/*
output:
Func C2
*/
```

Func匿名函
```go
package main

import "fmt"

func main() {
	c := func() { // 匿名函数
		fmt.Println("Func 匿名函数")
	}
	c()
}
```

闭包经典案例
```go
package main

import "fmt"

func main() {
	d := closure(10)
	fmt.Println(d(1))
	fmt.Println(d(2))
}

// closure 封闭包函数
func closure(x int) func(int) int {
	fmt.Printf("%p\n", &x)
	return func(y int) int {
		fmt.Printf("%p\n", &x)
		return x + y
	}
}
/*
output:
0xc4200160b0
0xc4200160b0
11
0xc4200160b0
12
*/
```

## defer

- defer的执行方式类似其它语言中的析构函数，在函数体执行结束后按照调用顺序的相反顺序逐个执行(先进后出，后进先出)
- 即使函数发生严重错误也会执行
- 支持匿名函数的调用
- 常用于资源清理、文件关闭、解锁以及记录时间等操作
- 通过与匿名函数配合可在return之后修改函数计算结果
- 如果函数体内某个变量作为defer时匿名函数的参数，则在定义defer时即已经获得了拷贝，否则则是引用某个变量的地址

- Go 没有异常机制，但有 panic/recover 模式来处理错误
- Panic 可以在任何地方引发，但recover只有在defer调用的函数中有效

defer 先进后出，后进先出
```go
package main

import "fmt"

func main() {
	fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")
}
/*
a
c
b
*/

```

案例2
```go
package main

import "fmt"

func main() {
	DeferA()
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
}
/*
0
1
2
3
4
----------------------------
4
3
2
1
0
*/
```

defer 在闭包的时候请多注意
```go
// DeferB  defer 在闭包的时候请多注意
func DeferB() {
	// fmt.Println("----------------------------")
	for i := 0; i < 5; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}
/*
5
5
5
5
5
*/
```

defer recover恢复测试
```go
package main

import "fmt"

func main() {
	DeferC()
	DeferC1()
	DeferC2()
}

func DeferC() {

	fmt.Println("func DeferC")
}
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
		第一种情况（没有上面的defer函数）output:
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

func DeferC2() {
	fmt.Println("func DeferC2")
}



```

## 课堂作业

- 运行以下程序并分析输出结果。
```go
package main

import "fmt"

func main() {
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
/*
output:
closure i =  4
closure i =  4
closure i =  4
closure i =  4
defer_closure i =  4
defer i =  3
defer_closure i =  4
defer i =  2
defer_closure i =  4
defer i =  1
defer_closure i =  4
defer i =  0
*/
```