package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup
	resultCh := make(chan int)

	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			resultCh <- n * n
		}(num)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	sum := 0
	for res := range resultCh {
		sum += res
	}

	fmt.Println("Sum: ", sum)
}
