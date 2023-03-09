package main

import (
	"fmt"
	"sync"
	"time"
)

// 3 способ - блокировка мьютекса, к которому обращается горутина
func main() {
	var stopMutex sync.Mutex
	a := true

	go func() {
		for {
			stopMutex.Lock()
			stopMutex.Unlock()

			fmt.Println("Goroutine doing smth...")
			time.Sleep(time.Millisecond * 2500)
		}
	}()

	// заблокируем мьютекс через 5 секунд
	time.AfterFunc(time.Second*5, func() {
		fmt.Println("Calling lock on mutex. Goroutine stopped")
		stopMutex.Lock()
	})

	// Бесконечный цикл чтоб мейн поток не завершался пока работает горутина
	for a {

	}
}
