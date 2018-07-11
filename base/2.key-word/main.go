package main

import (
	"fmt"
	"runtime"
	"time"

	bdchan "github.com/antyiwei/gobase/base/2.key-word/channel"
)

func main() {

	// TestIotaMain() // test iota main

	// TestInt2String() // test int to string
	ListenGorout()
	bdchan.TestChannel() // test channel

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
