package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type WaitGroup struct {
	counter int64
}

func (wg *WaitGroup) Add(n int64) {
	atomic.AddInt64(&wg.counter, n)
}

func (wg *WaitGroup) Done() {
	wg.Add(-1)

	if atomic.LoadInt64(&wg.counter) < 0 {
		panic("negative wait group counter")
	}
}

func (wg *WaitGroup) Wait() {
	for {
		if atomic.LoadInt64(&wg.counter) == 0 {
			return
		}
	}
}

func main() {
	var wg WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		time.Sleep(500 * time.Millisecond)
		fmt.Println("goroutine 1")
	}()

	go func() {
		defer wg.Done()
		time.Sleep(200 * time.Millisecond)
		fmt.Println("goroutine 2")
	}()

	wg.Wait()
	fmt.Println("all goroutines are done")
}
