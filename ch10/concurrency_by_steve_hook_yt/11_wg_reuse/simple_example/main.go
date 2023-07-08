package main

import (
	"sync"
	"time"
)

func main() {
	// calling operations on wg before even waits
	// i.e. reusing it -> cause panic
	// until it has executed wait
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(time.Second)
		wg.Done()
		wg.Add(1) // cause panic : WaitGroup is reused before previous Wait has returned
	}()
	wg.Wait()
}
