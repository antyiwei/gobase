# 逻辑判断

## 判断语句if

- 条件表达式没有括号
- 支持一个初始化表达式（可以是并行方式）
- 左大括号必须和条件语句或else在同一行
- 支持单行模式
- 初始化语句中的变量为block级别，同时隐藏外部同名变量
- 1.0.3版本中的编译器BUG

```go
package main

import "fmt"

/*
> Output:
command-line-arguments
小于等于6
1
true
> Elapsed: 0.317s
> Result: Success
 */
func main() {
    a := true
    if a, b, c := 1, 2, 3; a+b+c > 6 {
        fmt.Println("大于6")
    } else {
        fmt.Println("小于等于6")
        fmt.Println(a)
    }
    fmt.Println(a)
}
```

## 循环语句for

- Go只有for一个循环语句关键字，但支持3种形式
- 初始化和步进表达式可以是多个值
- 条件语句每次循环都会被重新检查，因此不建议在条件语句中使用函数，尽量提前计算好条件并以变量或常量代替
- 左大括号必须和条件语句在同一行

```go
package main

import "fmt"

func main() {
	a := 1
	for {
		a++
		if a > 3 {
			break
		}
	}
	fmt.Println(a)
}
```

```go
package main

import "fmt"

func mian() {
	a := 1
	for a <= 3 {
		a++
	}
	fmt.Println(a)
}
```

```go
package main

import "fmt"

func main() {
	a := 1
	for i := 0; i < 3; i++ {
		a++
	}
	fmt.Println(a)
}
```

## 选择语句switch

- 可以使用任何类型或表达式作为条件语句
- 不需要写break，一旦条件符合自动终止
- 如希望继续执行下一个case，需使用fallthrough语句
- 支持一个初始化表达式（可以是并行方式），右侧需跟分号
- 左大括号必须和条件语句在同一行

```go
package main

import "fmt"

func main() {
	switch a {
	a := 1
	case 0:
		fmt.Println("a=0")
	case 1:
		fmt.Println("a=1")
	}
	fmt.Println(a)
}
```

```go
package main

import "fmt"

func main() {
	a := 1
	switch {
	case a >= 0:
		fmt.Println("a=0")
	case a >= 1:
		fmt.Println("a=1")
	}
	fmt.Println(a)
}
```

```go
package main

import "fmt"

func main() {
	switch a := 1; {
	case a >= 0:
		fmt.Println("a = 0")
	case a >= 1:
		fmt.Println("a = 1")
	}
}
```

## 跳转语句goto, break, continue

- 三个语法都可以配合标签使用
- 标签名区分大小写，若不使用会造成编译错误
- Break与continue配合标签可用于多层循环的跳出
- Goto是调整执行位置，与其它2个语句配合标签的结果并不相同

```go
package main

import "fmt"

func main() {
LABEL:
	for i := 0; i < 10; i++ {
		if i > 2 {
			break LABEL
		} else {
			fmt.Println(i)
		}
	}

}
```

```go
package main

import "fmt"

func main() {
LABEL:
	for i := 0; i < 10; i++ {
		for {
			fmt.Println(i)
			continue LABEL
		}
	}
}
```

## 课堂作业

- 将下图中的continue替换成goto，程序运行的结果还一样吗？
- 请尝试并思考为什么。

```go
func GotoTest3() {
LABEL:
	for i := 0; i < 10; i++ {
		for {
			fmt.Println(i)
			goto LABEL
		}
	}
}
```
- 注意：会出现死循环，小心

## for 中的 break，continue, return
- break 跳出最近的一个for循环
- continue 跳出当前循环
- return 跳出当前程序，返回到最原始状态