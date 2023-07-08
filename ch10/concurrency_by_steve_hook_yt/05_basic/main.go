package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3) // we want to run 3 goroutines
	go func() {
		defer wg.Done()
		fmt.Println("1")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("2")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("3")
	}()

	wg.Wait() // It will block until all goroutines have called their done methods
}
