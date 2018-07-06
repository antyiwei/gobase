package main

import (
	"fmt"
	"time"

	"github.com/henrylee2cn/goutil/pool"
	"github.com/panjf2000/ants"
)

func main() {
	// TestGoPool()
	AntPoolsTest()
}

func TestGoPool() {
	goroupool := pool.NewGoPool(10000, 0)
	fmt.Println("开始时间：", time.Now().Format("2006-01-02 15:04:05.999"))
	retryTimes := 0
	var err error
	for i := 0; i < 10000001; i++ {
		err = goroupool.Go(func() {
			A(i)
		})
		if err != nil {
			retryTimes++
		}
	}
	fmt.Printf("retryTimes: %d \n", retryTimes)
	fmt.Println("开始时间：", time.Now().Format("2006-01-02 15:04:05.999"))
}
func FmtStr() {

}
func A(i int) {
	if i > 1000000 {
		fmt.Printf("my name is %d \n", (i + 1))
	}

	time.Sleep(time.Duration(10) * time.Microsecond)

	return
}

// func TestGoPool2() {
// 	gp := NewGoPool(10, 0)
// 	wg := new(sync.WaitGroup)
// 	retryTimes := 0
// 	var err error
// 	for i := 0; i < 100; i++ {
// 		wg.Add(1)
// 		a := i
// 		err = gp.Go(func() {
// 			t.Log("done:", a)
// 			wg.Done()
// 		})
// 		if err != nil {
// 			retryTimes++
// 			i--
// 			t.Log(err)
// 			wg.Done()
// 		}
// 	}
// 	wg.Wait()
// 	gp.Stop()
// 	t.Logf("retryTimes: %d", retryTimes)
// }

/*  ants goroutine pools */
func AntPoolsTest() {
	fmt.Println("开始时间：", time.Now().Format("2006-01-02 15:04:05.999"))
	// set 10000 the size of goroutine pool
	p, _ := ants.NewPool(100000)
	retryTimes := 0
	var err error
	for i := 0; i < 1000004; i++ {
		// submit a task
		p.Submit(func() error {
			A(i)
			return nil
		})

		if err != nil {
			retryTimes++
		}
	}
	fmt.Printf("retryTimes: %d \n", retryTimes)
	fmt.Println("开始时间：", time.Now().Format("2006-01-02 15:04:05.999"))
}
