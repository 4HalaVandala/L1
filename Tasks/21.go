package main

import "fmt"

/*
Паттерн "Адаптер" позволяет приспособить интерфейс одного класса к интерфейсу другого класса, чтобы они могли взаимодействовать.
*/

// Square Допустим, у нас есть структура Square с методом Area(), который вычисляет площадь квадрата:
type Square struct {
	Side float64
}

func (s *Square) Area() float64 {
	return s.Side * s.Side
}

// Shape Теперь нам нужно использовать этот класс вместе с другим классом, который ожидает объект с методом GetArea().
// Мы можем создать адаптер для Square, чтобы он соответствовал интерфейсу требуемого класса:
type Shape interface {
	GetArea() float64
}

type SquareAdapter struct {
	*Square
}

func (sa *SquareAdapter) GetArea() float64 {
	return sa.Area()
}

/*
В данном примере мы создаем новый тип SquareAdapter, который включает в себя указатель на Square.
Метод GetArea() возвращает площадь квадрата, используя метод Area() из Square.
Теперь мы можем использовать SquareAdapter вместо Square в любом месте, где требуется объект с методом GetArea():
*/

func PrintArea(s Shape) {
	fmt.Println("Area:", s.GetArea())
}

func main() {
	square := &Square{Side: 5}
	adapter := &SquareAdapter{Square: square}
	PrintArea(adapter)
}
