package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nilmap map[string]int  // map[keyType]valueType
	fmt.Println(nilmap == nil) // true
	// nilmap["hello"] = 2  --> panic

	map2 := map[string]int{} // empty map
	map2["hello"] = 2
	fmt.Println(map2)

	teams := map[string][]string{ // key: string, value: slice of strings
		"Orcas":   []string{"Fred", "Ralph", "Bijou"},
		"Lions":   []string{"Sarah", "Peter", "Billie"},
		"Kittens": []string{"Waldo", "Raul", "Ze"},
	}
	fmt.Println(teams)

	// maps using `make`
	ages := make(map[int][]string, 2) // len = 0
	ages[1] = []string{"a", "b"}
	ages[2] = []string{"e", "f"}
	fmt.Println(ages)

	// comma ok idiom
	v, ok := teams["abc"] // if value not present, then v=default value, ok=false
	fmt.Println(v, ok, reflect.TypeOf(v))
	v2, ok := teams["Orcas"]
	fmt.Println(v2, ok)

	// deleting from map
	delete(teams, "Orcas")
}
