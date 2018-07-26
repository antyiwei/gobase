package main

import (
	"fmt"
	"time"
)

// TestChannel TestChannel
func TestChannel() error {
	// fmt.Println(" ---------- test DoTimer method -----------")
	// DoTimer()

	// fmt.Println(" ---------- test DoTicker method -----------")
	// DoTicker()

	// fmt.Println(" ---------- test DoClose method -----------")
	// DoClose()

	fmt.Println(" ---------- test DoWork method -----------")
	DoWork()

	return nil
}

// DoTimer DoTimer
func DoTimer() {
	fmt.Println("test channel DoTimer   begin ===========")
	timer1 := time.NewTimer(time.Second * 2)
	<-timer1.C
	fmt.Println("Timer 1 expired")

	fmt.Println("test channel timer.Stop   begin ===========")
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stoped")
	}
	time.Sleep(2 * time.Second)
}

// DoTicker DoTicker
func DoTicker() {
	fmt.Println("test channel ticker   begin ===========")
	ticker := time.NewTicker(time.Millisecond * 200) //ticker每500毫秒触发一次
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at ", t)
		}
	}()

	time.Sleep(1 * time.Second)
}

// DoClose DoClose
func DoClose() {
	go func() {
		time.Sleep(time.Hour)
	}()
	c := make(chan int, 10)
	c <- 1
	c <- 2
	close(c)
	/*  第一种 */
	// c <- 3

	/*  第二种 */
	// fmt.Println(<-c) //1
	// fmt.Println(<-c) //2
	// fmt.Println(<-c) //0
	// fmt.Println(<-c) //0

	/*  第三种 但是如果通过range读取，channel关闭后for循环会跳出：*/
	// for i := range c {
	// 	fmt.Println(i)
	// }

	/* 第四种 通过i, ok := <-c可以查看Channel的状态，判断值是零值还是正常读取的值。*/
	// 这种情况，需要 关闭通道赋值
	i, ok := <-c
	fmt.Printf("%d, %t", i, ok) //0, false

}

// DoWork DoWork
func DoWork() {
	done := make(chan bool, 1)
	go worker(done)

	time.Sleep(time.Second * 2)
	fmt.Println(" my name is  antyiwei ")
	// 等待任务完成
	var result, ok = <-done
	fmt.Println(result, ok)

	// var result2, ok2 = <-done
	// fmt.Println(result2, ok2)
}

func worker(done chan bool) {
	time.Sleep(3 * time.Second)

	// 处理工作任务

	// 通知任务已经完成
	done <- true

}
