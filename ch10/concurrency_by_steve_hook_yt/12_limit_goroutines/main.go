package main

import (
	"fmt"
	"sync"
	"time"
)

type request func()

func main() {
	now := time.Now()
	requests := map[int]request{}
	for i := 0; i < 100; i++ {
		f := func(n int) request {
			return func() {
				time.Sleep(100 * time.Millisecond)
				fmt.Println("request", n)
			}
		}
		requests[i] = f(i)
	}

	var wg sync.WaitGroup
	max := 10 // max 10 requests can be made
	for i := 0; i < len(requests); i += max {
		for j := i; j < i+max; j++ {
			wg.Add(1)
			go func(r request) {
				defer wg.Done()
				r()
			}(requests[j])
			// wg.Wait()  --> this will make the program sequential
		}
		wg.Wait() // wait for all the concurrent goroutines to finish => faster
		fmt.Println(max, "requests processed")
	}
	fmt.Println("elapsed", time.Since(now))
}
