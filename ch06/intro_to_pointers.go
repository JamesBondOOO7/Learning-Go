package main

import "fmt"

func main() {
	var x int32 = 10
	var y bool = true
	pointerX := &x
	pointerY := &y
	var pointerZ *string
	fmt.Println(pointerX, pointerY, pointerZ)

	var a = new(int)
	fmt.Println(a == nil) // prints false
	fmt.Println(*a)       // prints 0

	var b string // for primitive types
	pt := &b
	fmt.Println(pt, len(*pt))
}
