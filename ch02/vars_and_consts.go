package main

import "fmt"

const a int64 = 44
const (
	idKey   = "id"
	nameKey = "name"
)

func main() {
	// variables
	var x int = 10
	x, y := 20, 30

	fmt.Println(x, y)

	// constants - immutable
	const b = "hello"
	fmt.Println(a, b, idKey, nameKey)
}
