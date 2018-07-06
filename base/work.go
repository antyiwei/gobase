package main

/*
课堂作业

        既然导入多个包时可以进行简写，那么声明多个 常量、全局变量
或一般类型（非接口、非结构）是否也可以用同样的方法呢？

*/

// 常量定义
const (
	PI     = 3.14
	const1 = "1"
	const2 = 2
	const3 = 3
)

//  全局变量的申明和赋值
var (
	name  = "gopher"
	name1 = "1"
	name2 = 2
	name3 = 3
)

// 一般类型声明
type (
	newType int
	type1   float64
	type2   string
	type3   byte
)
