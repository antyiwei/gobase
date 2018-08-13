package main

import (
	"fmt"

	"github.com/antyiwei/gobase/bd_mode/factory"
	"github.com/antyiwei/gobase/bd_mode/singleton"
)

func main() {
	//singletonMode()// 单例模式

	//factoryMode() // 工厂模式

	//factionOperationMode() // 工厂

	factroyGirlFriendMode() // 抽象工厂
}

func singletonMode() {
	{
		/* 单例 */
		singleton.New() // 非线程安全

		singleton.NewSafe() // 线程安全单例模式

	}
}

func factoryMode() {

	Operator := factory.NewOperateFactory().CreateOperate("+")
	fmt.Printf("add result is %d\n", Operator.Operate(1, 2))
}

func factionOperationMode() {
	fac := &(factory.AddFactory{})
	oper := fac.CreateOperation()
	oper.SetA(1)
	oper.SetB(2)
	fmt.Println(oper.GetResult())
}

func factroyGirlFriendMode() {
	a := factory.GetGirlFriend("Indian")
	fmt.Println()
}
