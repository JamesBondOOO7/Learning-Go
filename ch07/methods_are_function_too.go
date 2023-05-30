package main

import "fmt"

type Adder struct {
	start int
}

func (a Adder) AddTo(val int) int {
	return a.start + val
}

func main() {
	myAdder := Adder{start: 10}
	fmt.Println(myAdder.AddTo(5))

	// assigning method to a variable OR
	// pass it to a variable of the same type (here, func(int) int)
	// - METHOD VALUE
	f1 := myAdder.AddTo
	fmt.Println(f1(10))

	// create a func from the type itself
	// the first parameter is the receiver for the method
	// (our function signature is func(Adder, int) int)
	// - METHOD EXPRESSION
	f2 := Adder.AddTo
	fmt.Println(f2(myAdder, 15))
}
