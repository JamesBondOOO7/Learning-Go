package main

import (
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// in a loop, iterations happen very fast
	// it would happen that `wg.Add` would be executed for the next iteration
	// even before all 3 `wg.Done` and `wg.Wait` executes
	// => cause panic : WaitGroup is reused before previous Wait has returned
	for i := 0; i < 100; i++ {
		go func() {
			wg.Add(3)
			go func() {
				wg.Done()
			}()
			go func() {
				wg.Done()
			}()
			go func() {
				wg.Done()
			}()
			wg.Wait()
		}()
	}
}
