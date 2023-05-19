package main

import "fmt"

func main() {
	samples := [][]string{
		{
			"hello", "james",
		},
		{
			"hi", "ethan", "hunt",
		},
	}

outer:
	for _, s1 := range samples {
		fmt.Println("---")
		for _, s2 := range s1 {
			for i, r := range s2 {
				fmt.Println(i, r, string(r))
				if r == 'a' {
					continue outer
				}
			}
			fmt.Println(s2, "ends")
		}
	}

}
