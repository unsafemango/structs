package main

import "fmt"

func generic() {
	result := add(1, 2)
	fmt.Println(result)
}

// generic function to correctly infer the type of value we are working with
func add[T int | float64 | string](a, b T) T {
	return a + b
}
