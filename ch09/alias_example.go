package main

import "fmt"

type Foo struct {
	x int
	S string
}

func (f Foo) Hello() string {
	return "hello"
}

func (f Foo) goodBye() string {
	return "goodbye"
}

// If we want to allow users to access Foo by the name Bar

type Bar = Foo

func MakeBar() Bar {
	bar := Bar{
		x: 20,
		S: "Hello",
	}
	var f Foo = bar
	fmt.Println(f.Hello())
	return bar
}
