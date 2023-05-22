package main

import "fmt"

func failedUpdate(g *int) {
	x := 10
	g = &x
}
func main() {
	var f *int // f is nil
	failedUpdate(f)
	fmt.Println(f) // prints nil
}
