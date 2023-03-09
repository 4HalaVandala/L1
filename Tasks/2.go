package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	for _, num := range nums {
		wg.Add(1)

		go func(n int) {
			defer wg.Done()
			square := n * n
			fmt.Println(square)
		}(num)
	}
	wg.Wait()
}
