package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		v := 1
		fmt.Println("Hello1")
		ch1 <- v
		fmt.Println("Hello2")
		v2 := <-ch2
		fmt.Println("In anonymous goroutine", v, v2)
	}()
	time.Sleep(1 * time.Second)

	v := 2
	var v2 int
	select {
	case ch2 <- v:
		fmt.Println("Case1")
	case v2 = <-ch1:
		fmt.Println("Case2")
	}
	fmt.Println("In main goroutine", v, v2)
}
