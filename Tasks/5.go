package main

import (
	"fmt"
	"time"
)

func main() {
	dataChan := make(chan int)
	duration := 5 // время работы программы в секундах

	go func() {
		for i := 0; ; i++ {
			select {
			case <-time.After(time.Second * time.Duration(duration)):
				fmt.Println("Stopping data generation...")
				close(dataChan)
				return
			case dataChan <- i:
			}
		}
	}()

	for data := range dataChan {
		fmt.Println("Received:", data)
	}
	fmt.Println("Finished")
}
