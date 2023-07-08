package main

import (
	"fmt"
	"time"
)

func say(s string, t int) {
	for i := 0; i < 5; i++ {

		if t == -1 {
			time.Sleep(10 * time.Millisecond)
		} else {
			time.Sleep(1000 * time.Millisecond)
		}
		fmt.Println(s)
	}
}

func main() {

	// CASE 1
	//go say("world", 1)
	//// main goroutine will be executed, no time left for the above goroutine
	//say("hello", -1)

	// CASE 2
	go say("world", -1)
	// main goroutine has to wait (goes to sleep -> blocked!), hence context switches to the above goroutine
	// It is really fast and executes the work
	say("hello", 1)
}
