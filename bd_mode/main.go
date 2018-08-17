package main

import (
	"github.com/antyiwei/gobase/bd_mode/singleton"
)

func main() {
	//singletonMode()// 单例模式

	factroyGirlFriendMode() // 抽象工厂
}

func singletonMode() {
	{
		/* 单例 */
		singleton.New() // 非线程安全

		singleton.NewSafe() // 线程安全单例模式

	}
}
