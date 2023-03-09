package main

import (
	"fmt"
	"time"
)

// Первый способ - передача каналу сигнала об остановке

func main() {
	stopChan := make(chan struct{})
	a := true

	// Запускаем горутину
	go func() {
		for {
			select {
			case <-stopChan:
				fmt.Println("Stopping goroutine...")
				return
			default:
				fmt.Println("Goroutine doing smth...")
				time.Sleep(time.Millisecond * 2500)
			}
		}
	}()

	// отправляем сигнал об остановке через 5 секунд
	time.AfterFunc(time.Second*5, func() {
		close(stopChan)
	})

	// Бесконечный цикл чтоб мейн поток не завершался пока работает горутина
	for a {

	}
}
