package main

import (
	"fmt"
	"math"
)

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// создаем карту для группировки значений по диапазону температур
	groups := make(map[int][]float64)

	// проходим по всем значениям и добавляем их в соответствующую группу
	for _, temp := range temps {
		group := int(math.Floor(temp/10)) * 10
		groups[group] = append(groups[group], temp)
	}

	// выводим все группы на экран
	for group, values := range groups {
		fmt.Printf("%d: %v\n", group, values)
	}
}
