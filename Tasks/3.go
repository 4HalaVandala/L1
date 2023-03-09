package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	nums := []int{2, 4, 6, 8, 10}
	results := make(chan int, len(nums))

	// создаем пул горутин
	pool := sync.Pool{
		New: func() interface{} {
			return make([]byte, 4)
		},
	}

	for _, n := range nums {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			buf := pool.Get().([]byte)
			defer pool.Put(buf)
			res := x * x
			str := fmt.Sprintf("%d\n", res)
			copy(buf, str)
			results <- res
		}(n)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	sum := 0
	for res := range results {
		sum += res
	}

	fmt.Println(sum)
}
