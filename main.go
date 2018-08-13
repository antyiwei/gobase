package main

import (
	"fmt"
	"runtime"

	"github.com/shopspring/decimal"
)

func main() {
	fmt.Println("Hello antyiwei world!!!")
	var sb float64
	sb = 0.1
	fmt.Println(sb)
	var a, b float32
	a = 8.5
	b = 4.64
	fmt.Println(a, b)
	fmt.Println(float32(a + b))

	c := decimal.NewFromFloat(8.5)
	d := decimal.NewFromFloat(4.64)
	f := decimal.Sum(c, d)
	fmt.Println(decimal.Sum(c, d).StringFixed(2))

	fmt.Println(f.Float64())

	var n float64 = 0
	for i := 0; i < 1000; i++ {
		n += .01
	}
	fmt.Println(n)

	{
		var n float64 = 0
		new := decimal.NewFromFloat(n)
		add := decimal.NewFromFloat(0.01)
		for i := 0; i < 1000; i++ {
			new = new.Add(add)
		}
		fmt.Println(new)
		fmt.Println(new.StringFixed(20))
		fmt.Println(new.Float64())
		fmt.Println(add.Float64())
	}

	//	 打印当前系统
	fmt.Println(runtime.GOOS)
}
