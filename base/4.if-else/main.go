package main

func main() {

	// 测试for
	ForTest1() // 第一个
	ForTest2() // 第二个
	ForTest3() // 第三个

	// 测试if
	IfTest1() // 第一个

	// 测试 switch
	SwitchTest1() // SwitchTest1 熟悉switch可以使用任何类型或表达式作为条件语句不需要写break，一旦条件符合自动终止
	SwitchTest2() // SwitchTest2 支持一个初始化表达式（可以是并行方式），右侧需跟分号 左大括号必须和条件语句在同一行
	SwitchTest3() // SwitchTest3 如希望继续执行下一个case，需使用fallthrough语句
}
