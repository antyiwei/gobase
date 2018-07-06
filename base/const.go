package main

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
