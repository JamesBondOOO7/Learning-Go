package main

import "fmt"

type person struct {
	name string
	age  int
	pet  string
}

func main() {
	p1 := person{"John", 52, "sheero"}
	fmt.Println(p1)

	p2 := person{
		age:  30,
		name: "Beth",
	}
	fmt.Println(p2)

	// anonymous structs
	pet := struct {
		name string
		kind string
	}{
		name: "Fido",
		kind: "dog",
	}
	fmt.Println(pet)

	type firstPerson struct {
		name string
		age  int
	}

	type secondPerson struct {
		name string
		age  int
	}

	a1 := firstPerson{
		name: "James",
		age:  44,
	}

	a2 := secondPerson{
		name: "James",
		age:  44,
	}

	// both the structs should have the same names, order, and types
	fmt.Println(a1 == firstPerson(a2)) // true

	// anonymous struct
	var anonymousStudent struct {
		name string
		age  int
	}

	anonymousStudent = a1
	fmt.Println(a1 == anonymousStudent)
}
