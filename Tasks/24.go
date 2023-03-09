package main

import (
	"fmt"
	"math"
)

/*
В этом примере мы определяем структуру Point с полями x и y, а также конструктор NewPoint для создания новых экземпляров структуры.
Функция Distance принимает две точки и использует формулу расстояния между двумя точками на плоскости, чтобы вычислить расстояние между ними.
В функции main мы создаем две точки и вызываем функцию Distance для вычисления расстояния между ними.

Формула для нахождения расстояния между двумя точками на плоскости: AB = (sqrt) ((Xb-Xa)^2 + (Yb-Ya)^2)
*/

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func Distance(p1, p2 *Point) float64 {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	p1 := NewPoint(1.0, 2.0)
	p2 := NewPoint(4.0, 6.0)
	dist := Distance(p1, p2)
	fmt.Println(dist)
}
