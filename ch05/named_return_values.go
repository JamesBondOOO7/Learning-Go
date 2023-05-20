package main

import (
	"errors"
	"fmt"
	"os"
)

func divAndRemainderNamed(numerator int, denominator int) (result int, remainder int,
	err error) {
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return result, remainder, err
	}
	result, remainder = numerator/denominator, numerator%denominator
	return result, remainder, err
}

func divAndRemainderNamedV2(numerator int, denominator int) (result int, remainder int,
	err error) {

	// assign some values
	result, remainder = 20, 30

	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, errors.New("cannot divide by zero")
	}
	return numerator / denominator, numerator % denominator, nil
}

func main() {
	result, remainder, err := divAndRemainderNamed(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, remainder)

	result, remainder, err = divAndRemainderNamedV2(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, remainder)
}
