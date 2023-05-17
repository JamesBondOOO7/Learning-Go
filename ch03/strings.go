package main

import "fmt"

func main() {

	var s string = "Hello 🌞"
	fmt.Println(len(s)) // 10 and not 7!

	// rune to string
	var a rune = 'x'
	s = string(a)

	var num int = 78
	s = string(num)
	fmt.Println(s)
}
