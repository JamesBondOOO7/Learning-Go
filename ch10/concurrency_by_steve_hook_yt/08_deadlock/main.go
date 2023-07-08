package main

import "sync"

func main() {
	var wg sync.WaitGroup
	wg.Add(1) // adding 1 operation
	// but never called done
	wg.Wait() // block infinitely => deadlock
}
