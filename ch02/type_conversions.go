package main

import "fmt"

func main() {
	var x int = 10
	var y float64 = 30.8
	var z float64 = float64(x) + y
	var d int = x + int(y) // int trims the decimal part
	fmt.Println("z:", z, "d:", d)
}
