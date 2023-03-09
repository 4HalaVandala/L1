package main

import "fmt"

func main() {
	//индекс удаляемого элемента
	i := 1
	// создаем слайс чисел
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Array before deleting element: ", numbers)
	fmt.Println("Index element to delete: ", i)

	// удаляем 3-й элемент (индекс 2)
	index := i
	copy(numbers[index:], numbers[index+1:])
	numbers = numbers[:len(numbers)-1]

	fmt.Println("Array after: ", numbers) // [1 2 4 5]
}
