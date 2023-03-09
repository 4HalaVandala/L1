package main

import "fmt"

func setBit(num int64, pos uint, val bool) int64 {
	if val {
		// установить i-й бит в 1
		return num | (1 << pos)
	} else {
		// установить i-й бит в 0
		return num &^ (1 << pos)
	}
}

func main() {
	var num int64 = 10 // число 1010 в двоичной системе
	fmt.Printf("num = %d (%b)\n", num, num)

	// установить 2-й бит в 1
	num = setBit(num, 2, true)
	fmt.Printf("num = %d (%b)\n", num, num)

	// установить 3-й бит в 0
	num = setBit(num, 3, false)
	fmt.Printf("num = %d (%b)\n", num, num)
}
