# Go 编程基础

- 什么是Go？
    - Go是一门"并发支持、垃圾回收的编译型"系统编程语言，旨在创造一门具有在静态编译语言的高性能和动态语言的高效开发之间拥有良好平衡点的一门编程语言。
- Go的主要特点有哪些？
    - 类型安全和内存安全
    - 以非常直观和极低代价的方案实现 高并发
    - 高效的垃圾回收机制
    - 快速编译（同时解决C语言中头文件太多的问题）
    - 为多核计算机提供性能提升的方案
    - UTF-8编码支持

## Go内置关键字（25个均为小写）

```go
break       default           func        interface         select
case        defer             go          map               struct
chan        else              goto        package           switch
const       fallthrough       if          range             type
continue    for               import      return            var
```

- Go注释方法
    -   //：单行注释
    -   /* */：多行注释


## Go程序的一般结构：basic_structure.go

- Go程序是通过 package 来组织的（与python类似）只有 package 名称为 main 的包可以包含 main 函数一个可执行程序 有且仅有 一个 main 包
- 通过 import 关键字来导入其它非 main 包
- 通过 const 关键字来进行常量的定义
- 通过在函数体外部使用 var 关键字来进行全局变量的声明与赋值
- 通过 type 关键字来进行结构(struct)或接口(interface)的声明
- 通过 func 关键字来进行函数的声明

## 可见性规则

- Go语言中，使用 大小写 来决定该 常量、变量、类型、接口、结构或函数 是否可以被外部包所调用：根据约定，函数名首字母 小写 即为private

## Go基本类型

- 布尔型：bool
    - 长度：1字节
    - 取值范围：true, false
    - 注意事项：不可以用数字代表true或false

- 整型：int/uint
    - 根据运行平台可能为32或64位

- 8位整型：int8/uint8
    - 长度：1字节
    - 取值范围：-128~127/0~255
- 字节型：byte（uint8别名）


- 16位整型：int16/uint16
    - 长度：2字节
    - 取值范围：-32768~32767/0~65535
- 32位整型：int32（rune）/uint32
    - 长度：4字节
    - 取值范围：-2^32/2~2^32/2-1/0~2^32-1
- 64位整型：int64/uint64
    - 长度：8字节
    - 取值范围：-2^64/2~2^64/2-1/0~2^64-1
- 浮点型：float32/float64
    - 长度：4/8字节
    - 小数位：精确到7/15小数位


- 复数：complex64/complex128
    - 长度：8/16字节
- 足够保存指针的 32 位或 64 位整数型：uintptr

- 其它值类型：
    - array、struct、string
- 引用类型：
    - slice、map、chan

- 接口类型：inteface
- 函数类型：func


## 类型零值
- 零值并不等于空值，而是当变量被声明为某种类型后的默认值，通常情况下值类型的默认值为0，bool为false，string为空字符串
- 类型别名

```go
type(
    byte int8
    rune int32
    文本 string
)

var b 文本
b =  "中文啊亲"

```

## 单个变量的声明与赋值

- 变量的声明格式：var <变量名称> <变量类型>
- 变量的赋值格式：<变量名称> = <表达式>
- 声明的同时赋值：var <变量名称> [变量类型] = <表达式>

```go
    var a int // 变量的声明
    a  = 123 // 变量的赋值

    // 变量声明同时赋值
    var b int =  321
    // 上行的格式可以省略类型，由系统推断
    var c = 321
    // 变量声明与赋值的最简洁写法
    d:= 345

```

## 多个变量的声明与赋值

- 全局变量的声明可使用 var() 的方式进行简写
- 全局变量的声明不可以省略 var，但可使用并行方式
- 所有变量都可以使用类型推断
- 局部变量不可以使用 var() 的方式简写，只能使用并行方式

```go
var (
    // 使用常规方式
    aaa = "hello"
    // 使用并行方式以及类型推断
    sss,bbb = 1,2
    // ccc:= 3 // 不可以省略 var
)

    //多个变量的声明
    var a, b, c, d int
    // 多个变量的赋值
    a, b, c, d = 1, 2, 3, 4

    // 多个变量声明的同时赋值
    var e, f, g, h int = 5, 6, 7, 8
    // 省略变量类型，由系统推断
    var i, j, k, l = 9, 10, 11, 12
    // 多个变量声明与赋值的最简写法
    i, m, n, o := 12, 23, 434, 45

```

## 变量的类型转换

- Go中不存在隐式转换，所有类型转换必须显式声明
- 转换只能发生在两种相互兼容的类型之间
- 类型转换的格式：

```html
    <ValueA> [:]= <TypeOfValueA>(<ValueB>)
```

```go

    // 在相互兼容的两种类型之间进行转换
    var a float64 = 1.20
    b := int(a)

    // 以下表达式无法通过编译
    var c bool = true
    d := int(c)
```

## 课堂作业

请尝试运行以下代码，看会发生什么，并思考为什么。

```go
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
```

string() 表示将数据转换成文本格式，因为计算机中存储的任何东西
本质上都是数字，因此此函数自然地认为我们需要的是用数字65表示
的文本 A。

## 常量的定义

- 常量的值在编译时就已经确定
- 常量的定义格式与变量基本相同
- 等号右侧必须是常量或者常量表达式
- 常量表达式中的函数必须是内置函数

```go
// 定义单个常量
const a int = 1
const b = 'A'
const (
    text   = "123"
    length = len(test)
    num    = b * 20
)

// 同时定义多个变量
const i, j, k = 1, "2", '3'
const (
    text2, length2, num2 = "234", len(text2), k * 10
)
```

## 常量的初始化规则与枚举

- 在定义常量组时，如果不提供初始值，则表示将使用上行的表达式
- 使用相同的表达式不代表具有相同的值
- iota是常量的计数器，从0开始，组中每定义1个常量自动递增1
- 通过初始化规则与iota可以达到枚举的效果
- 每遇到一个const关键字，iota就会重置为0

```go
const (
    // a 与b 均为“A”
    a = "A"
    b
    c = iota
    d // d的值为3
)

const (
    e = iota
    f // f的值为1
)

// 星期枚举
const (
    // 第一个常量不能省略表达式
    Monday = iota
    Tuesday
    Wednesday
    Thursday
    Firday
    Saturday
    Sunday
)
```

## 运算符

- Go中的运算符均是从左至右结合
    优先级（从高到低）

```go
.   ^      !                                               （一元运算符）
.   *      /     %      <<    >>    &   &^ 
.   +      -     |      ^                                （二元运算符）
.   ==    !=     <      <=    >=    >
.   <-                                              （专门用于channel）
.   &&
.   ||
```


## 课堂作业

- 请尝试结合常量的iota与<<运算符实现计算机储存单位的枚举

```go
package main

import (
    "fmt"
    "reflect"
)

/*
fmt.Println(1 << (2 * 10
iota是golang语言的常量计数器,只能在常量的表达式中使用。
iota在const关键字出现时将被重置为0(const内部的第一行之前)，const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。
使用iota能简化定义，在定义枚举时很有用。
*/

type ByteSize float64

const (
    _           = iota
    KB ByteSize = (1 << (iota * 10))
    MB          // 1 << (10*2)
    GB          // 1 << (10*3)
    TB          // 1 << (10*4)
    PB          // 1 << (10*5)
    EB          // 1 << (10*6)
    ZB          // 1 << (10*7)
    YB          // 1 << (10*8)
)

func main() {
    fmt.Println(KB)
    fmt.Println(MB)
    fmt.Println(GB)
    fmt.Println(reflect.TypeOf(KB))
}
```

## 指针

 Go虽然保留了指针，但与其它编程语言不同的是，在Go当中不支持指针运算以及"->"运算符，而直接采用".选择符来操作指针目标对象的成员
- 操作符”&”取变量地址，使用”*”通过指针间接访问目标对象
- 默认值为 nil 而非 NULL

递增递减语句

    在Go当中，++ 与 -- 是作为语句而并不是作为表达式

