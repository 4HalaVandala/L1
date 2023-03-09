package main

import (
	"context"
	"fmt"
	"time"
)

// 2 способ - отмена контекста
func main() {
	ctx, cancel := context.WithCancel(context.Background())

	a := true

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Stopping goroutine...")
				return
			default:
				fmt.Println("Goroutine doing smth...")
				time.Sleep(time.Millisecond * 2500)
			}
		}
	}(ctx)

	// отправляем сигнал об отмене контекста через 5 секунд
	time.AfterFunc(time.Second*5, func() {
		cancel()
	})

	// Бесконечный цикл чтоб поток мейн не завершился
	for a {

	}
}
