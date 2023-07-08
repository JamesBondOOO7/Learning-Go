package main

import (
	"fmt"
	"runtime"
	"time"
)

type searchFunc func(string) []string

func searchData(s string, searchers []searchFunc) []string {

	result := make(chan []string)
	done := make(chan struct{})
	//defer close(done)

	// by this logic, the fastest search's result will be used
	for _, searcher := range searchers {
		go func(searcher searchFunc) {
			select {
			case result <- searcher(s):
				fmt.Println("I am running", <-result)
			case <-done:
				fmt.Println("X X X Done X X X")
			}
		}(searcher)
	}
	r := <-result
	fmt.Println("Oye result:", r)
	return r
}

func main() {

	// Capture starting number of goroutines.
	startingGs := runtime.NumGoroutine()
	startTime := time.Now()

	funcs := []searchFunc{
		func(s string) []string {
			fmt.Println("James Bond")
			time.Sleep(10 * time.Millisecond)
			return []string{"James Bond"}
		},
		func(s string) []string {
			fmt.Println("Hello World")
			time.Sleep(time.Second)
			return []string{"Hello World"}
		},
		func(s string) []string {
			fmt.Println("Mission Impossible")
			time.Sleep(2 * time.Second)
			return []string{"Mission Impossible"}
		},
		func(s string) []string {
			fmt.Println("Concurrency Concepts are fun :)")
			time.Sleep(3 * time.Second)
			return []string{"Concurrency Concepts are fun :)"}
		},
	}

	result := searchData("Test string", funcs)

	// Hold the program from terminating for 3 seconds to see
	// if any goroutines created by process terminate.
	time.Sleep(3 * time.Second)

	fmt.Println(">> Output:", result[0])

	// Capture ending number of goroutines.
	endingGs := runtime.NumGoroutine()

	endTime := time.Now()

	// Report the results.
	fmt.Println("========================================")
	fmt.Println("Number of goroutines before:", startingGs)
	fmt.Println("Number of goroutines after :", endingGs)
	fmt.Println("Number of goroutines leaked:", endingGs-startingGs)
	fmt.Println("Time taken", endTime.Sub(startTime))

}
