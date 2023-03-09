package main

import (
	"fmt"
	"sort"
)

/*
Функиця sort.Slice() позволяет сортироваться срезы по заданному критерию
С помощью функции-компаратора мы задаем порядок элементов
После вызова sort.Slice() массив arr будет отсортирован по возрастанию
*/

func main() {
	arr := []int{5, 2, 8, 3, 1, 6, 9, 4, 7}
	fmt.Println("Before sorting:", arr)

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	fmt.Println("After sorting:", arr)
}
