package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}

	// создаем каналы для передачи данных
	ch1 := make(chan int)
	ch2 := make(chan int)

	// запускаем первую горутину для записи чисел в канал ch1
	go func() {
		for _, num := range nums {
			ch1 <- num
		}
		close(ch1) // закрываем канал после записи всех чисел
	}()

	// запускаем вторую горутину для чтения из канала ch1 и записи в канал ch2
	go func() {
		for num := range ch1 {
			ch2 <- num * 2
		}
		close(ch2) // закрываем канал после записи всех результатов
	}()

	// читаем из канала ch2 и выводим результаты в stdout
	for res := range ch2 {
		fmt.Println(res)
	}
}
