package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Stopping goroutine...")
				return
			default:

				for i := 0; ; i++ {
					fmt.Println(i)
				}
			}
		}
	}(ctx)

	// отправляем сигнал об отмене контекста через 5 секунд
	time.AfterFunc(time.Second*3, func() {
		cancel()
	})
}
