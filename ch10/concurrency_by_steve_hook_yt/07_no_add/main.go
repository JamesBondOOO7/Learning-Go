package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// not adding this goroutine => not waiting for it
	go func() {
		defer wg.Done()
		time.Sleep(300 * time.Millisecond)
		fmt.Println("goroutine done")
	}()
	wg.Wait() // it executes immediately
	fmt.Println("executed immediately")
}
