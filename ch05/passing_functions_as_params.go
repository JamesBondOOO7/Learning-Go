package main

import (
	"fmt"
	"sort"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	people := []Person{
		{"Tony", "Stark", 55},
		{"James", "Bond", 54},
		{"Ethan", "Hunt", 51},
	}
	fmt.Println(people)

	// Sorting by LastName
	sort.Slice(people, func(i int, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	fmt.Println(people)

	// Sorting by Age
	sort.Slice(people, func(i int, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println(people)
}
