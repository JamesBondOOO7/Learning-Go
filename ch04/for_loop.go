package main

import "fmt"

func main() {

	completeForStatement()
	conditionOnlyForStatement()
	forRange()
}

func forRange() {

	evenVals := []int{2, 4, 6, 8, 10}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}

	// don't want to access index/keys
	for _, v := range evenVals {
		fmt.Println(v)
	}

	// just want the key, not the value
	uniqueNames := map[string]bool{"Fred": true, "Raul": true, "Wilma": true}
	for k := range uniqueNames {
		fmt.Println(k)
	}

	// Iterating over strings
	samples := []string{"hello", "apple_π!"}
	for _, sample := range samples {
		for i, r := range sample {
			// index, int value, its conversion to string
			fmt.Println(i, r, string(r))
		}
		fmt.Println()
	}

	// for-range loop is a copy
	for _, v := range evenVals {
		v *= 2
	}
	fmt.Println(evenVals)

}

func conditionOnlyForStatement() {
	i := 1
	for i < 100 {
		fmt.Println(i)
		i *= 2
	}
}

func completeForStatement() {
	// var is not legal here, always use `:=`
	// you can shadow a variable here
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
