package main

import "fmt"

func main() {
	set1 := []int{88, 33, 11, 1, 2, 3, 4, 5}
	set2 := []int{9, 10, 3, 4, 5, 6, 7}

	intersection := make([]int, 0)

	for _, value1 := range set1 {
		for _, value2 := range set2 {
			if value1 == value2 {
				intersection = append(intersection, value1)
			}
		}
	}

	fmt.Println(intersection)
}
