package main

import (
	"fmt"
	"time"
)

func main() {
	dataChan := make(chan int)
	t := 2
	duration := time.Second * time.Duration(t)
	timer := time.NewTimer(time.Duration(duration))

	go func() {
		for i := 0; ; i++ {
			select {
			case <-timer.C:
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
