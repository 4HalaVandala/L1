package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			fmt.Println(i, i*i)
			m.Store(i, i*i)
		}(i)
	}

	wg.Wait()
}
