package main

import "fmt"

func main() {
	/*     {
	       var wg sync.WaitGroup
	       var count int
	       var ch = make(chan bool, 1)
	       for i := 0; i < 10; i++ {
	           wg.Add(1)
	           go func() {
	               ch <- true
	               count++
	               time.Sleep(time.Millisecond)
	               count--
	               <-ch
	               wg.Done()
	           }()
	       }
	       wg.Wait()
	   } */
	/* 	{

			HappensBefore()
	    } */

	{
		var v T
		v.f()
		return
	}
}

/*  */
var c = make(chan int)
var a string

func HappensBefore() {
	go f()
	<-c
	fmt.Println("mama mi ya")
	print(a)
}
func f() {
	a = "Hello,world" // (1)
	c <- 0            //(2)
}

/*  */
type T int

func (t T) f() {
	fmt.Println("hello world!")
}
