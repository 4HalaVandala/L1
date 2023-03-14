package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"
)

func generateData() int {
	return rand.Intn(100)
}

func worker(id int, dataChan chan int) {
	for {
		select {
		case data := <-dataChan:
			fmt.Printf("Worker %d: %d\n", id, data)
		}
	}
}

func main() {
	dataChan := make(chan int)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	_, cancel := context.WithCancel(context.Background())

	numWorkers := 5

	for i := 0; i < numWorkers; i++ {
		go worker(i, dataChan)
	}

	for {
		select {
		case <-sigChan:
			fmt.Println("Stopping program...")
			cancel()
			close(dataChan)
			return
		default:
			data := generateData()
			dataChan <- data
			time.Sleep(time.Second)
		}
	}
}
