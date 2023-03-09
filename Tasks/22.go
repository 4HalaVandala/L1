package main

import (
	"fmt"
	"math/big"
)

/*
В этой программе мы используем пакет "math/big", который предоставляет функциональность для работы с большими числами.
Мы создаем две переменные a и b с помощью функции NewInt(), которая принимает значение типа int64 и возвращает указатель на объект типа big.Int.
*/

func main() {
	a := big.NewInt(100000000) // задаем значение переменной a
	b := big.NewInt(200000000) // задаем значение переменной b

	// умножение a на b
	mul := new(big.Int)
	mul.Mul(a, b)
	fmt.Println("Multiplication result:", mul)

	// деление a на b
	div := new(big.Int)
	div.Div(a, b)
	fmt.Println("Division result:", div)

	// сложение a и b
	sum := new(big.Int)
	sum.Add(a, b)
	fmt.Println("Sum result:", sum)

	// вычитание b из a
	sub := new(big.Int)
	sub.Sub(a, b)
	fmt.Println("Subtraction result:", sub)
}
