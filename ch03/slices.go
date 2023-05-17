package main

import "fmt"

func main() {

	var x = []int{10, 20, 30}
	fmt.Println(x)

	var y = []int{1, 5: 4, 6, 10: 100, 15} // sparse
	fmt.Println(y)

	var z []int
	fmt.Println(z == nil) // true

	fmt.Println(len(x))

	z = append(z, 7, 8, 9)
	fmt.Println(z)

	// len vs cap
	fmt.Println(x, len(x), cap(x))
	x = append(x, 40)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 50)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 60)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 70)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 80)

	q := make([]int, 5) // len = cap = 5, elements init to 0
	fmt.Println(q, len(q), cap(q))
	q = append(q, 10)
	fmt.Println(q, len(q), cap(q))

	r := make([]int, 5, 10) // len = 5, cap = 10
	fmt.Println(r, len(r), cap(r))

	s := make([]int, 0, 10) // len = 0, cap = 10
	s = append(s, 5, 6, 7, 8)
	fmt.Println(r, len(s), cap(s))

	// Slicing a slice
	l := []int{1, 2, 3, 4}
	m := l[:2]
	n := l[1:]
	l[1] = 20
	m[0] = 10
	n[1] = 30
	fmt.Println("l:", l)
	fmt.Println("m:", m)
	fmt.Println("n:", n)

	// array to slice
	arr := []int{1, 2, 3, 4}
	brr := make([]int, 2)
	num := copy(brr, arr) // len is considered, not cap
	fmt.Println(brr, num)

	num = copy(arr[:3], arr[1:])
	fmt.Println(arr, num)
}
