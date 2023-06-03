package main

import (
	"ch01/hello/ch09/package_example/formatter"
	"ch01/hello/ch09/package_example/math"
	"fmt"
)

func main() {
	num := math.Double(2)
	output := print.Format(num)
	fmt.Println(output)
}
