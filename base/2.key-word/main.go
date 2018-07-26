package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	// TestIotaMain() // test iota main

	// TestInt2String() // test int to string
	// ListenGorout()
	// bdchan.TestChannel() // test channel

	// learning goroutine and channel
	time.Sleep(60 * time.Second)
}

func ListenGorout() {
	var goroutineNum int
	ticker := time.NewTicker(time.Second * 1)
	go func() {
		for t := range ticker.C {
			goroutineNum = runtime.NumGoroutine()
			fmt.Println("Tick at ", t, "goroutine num:", goroutineNum)
		}
	}()
}
