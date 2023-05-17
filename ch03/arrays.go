package main

import "fmt"

func main() {
	var x [3]int
	fmt.Println(x)

	var y = [3]int{10, 20, 30}
	fmt.Println(y)

	// sparse array : [1, 0, 0, 0, 0, 4, 6, 0, 0, 0, 100, 15]
	var z = [12]int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println(z)

	var y2 = [...]int{10, 20, 30}
	fmt.Println(y == y2) // true

	const s int = 5
	var arr [s]int
	fmt.Println(arr)
}
