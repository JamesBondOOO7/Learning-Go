package main

import "sync"

func main() {
	var wg sync.WaitGroup
	// not adding any operation => WaitGroup counter = 0
	wg.Done() // but calling done -> panic
	// negative WaitGroup counter (i.e. = -1)
}
