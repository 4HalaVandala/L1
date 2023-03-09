package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	// Количество воркеров
	numWorkers := 5

	// Канал для записи данных
	dataChan := make(chan string)

	// Запуск воркеров
	wg := &sync.WaitGroup{}
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(dataChan, wg)
	}

	// Graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		fmt.Println("Received signal, shutting down...")
		close(dataChan)
	}()

	// Запись данных в канал
	for i := 0; ; i++ {
		select {
		case <-sigChan:
			fmt.Println("Stopping data generation...")
			return
		default:
			dataChan <- fmt.Sprintf("Data %d", i)
		}
	}

	// Ожидание завершения воркеров
	wg.Wait()
	fmt.Println("All workers stopped")
}

func worker(dataChan chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for data := range dataChan {
		fmt.Println(data)
	}
}
