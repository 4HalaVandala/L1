package main

import "fmt"

func main() {
	// Создаем список строк для обработки
	myList := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создаем map с уникальными элементами в качестве ключей и значением true
	uniqueMap := make(map[string]bool)

	// Проходим по списку строк и добавляем уникальные элементы в map
	for _, element := range myList {
		if _, ok := uniqueMap[element]; !ok {
			uniqueMap[element] = true
		}
	}

	// Выводим результаты
	fmt.Println("Исходный список:", myList)
	fmt.Println("Структура данных с уникальными элементами:", uniqueMap)
}
