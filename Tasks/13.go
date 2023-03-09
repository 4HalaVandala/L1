package main

import "fmt"

func main() {
	a, b := 5, 10
	fmt.Println("Before: a =", a, ", b =", b)

	a = a + b
	b = a - b
	a = a - b

	fmt.Println("After: a =", a, ", b =", b)
}
