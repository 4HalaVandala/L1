package main

import (
	"fmt"
	"sort"
)

/*
В Go встроенный метод бинарного поиска - это функция sort.Search(), которая ищет элемент в отсортированном срезе и возвращает его индекс.
Для реализации бинарного поиска мы должны сначала отсортировать срез, а затем вызвать функцию sort.Search() с нужными аргументами.
*/

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("Array:", arr)
	fmt.Println("Index of 5:", sort.SearchInts(arr, 5))
	fmt.Println("Index of 10:", sort.SearchInts(arr, 10))
}
