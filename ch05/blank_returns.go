package main

import (
	"errors"
	"fmt"
)

func divAndRemainderBlank(numerator, denominator int) (result int, remainder int,
	err error) {
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return
	}
	result, remainder = numerator/denominator, numerator%denominator
	return
}

func main() {
	result, remainder, err := divAndRemainderBlank(5, 2)
	fmt.Println(result, remainder, err)
}
