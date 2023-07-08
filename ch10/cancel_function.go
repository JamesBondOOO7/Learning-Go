package main

import (
	"fmt"
	"runtime"
	"time"
)

func countTo(max int) (<-chan int, func()) {
	ch := make(chan int)
	done := make(chan struct{})
	cancel := func() {
		close(done)
	}
	go func() {
		for i := 0; i < max; i++ {
			select {
			case <-done:
				return
			case ch <- i:
			}
		}
		close(ch)
	}()
	return ch, cancel
}
func main() {
	startGs := runtime.NumGoroutine()
	ch, cancel := countTo(10)
	for i := range ch {
		if i > 5 {
			break
		}
		fmt.Println(i)
	}
	cancel()
	time.Sleep(time.Second)
	endGs := runtime.NumGoroutine()
	fmt.Println("Goroutines Leaked:", endGs-startGs)
}
