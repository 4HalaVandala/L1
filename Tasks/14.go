package main

import (
	"fmt"
	"reflect"
)

func main() {
	//Switch type подходит для всего кроме канала
	var i interface{} = 10
	switch i.(type) {
	case int:
		fmt.Println("Переменная i имеет тип int")
	case string:
		fmt.Println("Переменная i имеет тип string")
	case bool:
		fmt.Println("Переменная i имеет тип bool")
	case chan int:
		fmt.Println("Переменная i имеет тип chan int")
	default:
		fmt.Println("Неизвестный тип переменной i")
	}

	var s interface{} = "Hello"
	switch s.(type) {
	case int:
		fmt.Println("Переменная s имеет тип int")
	case string:
		fmt.Println("Переменная s имеет тип string")
	case bool:
		fmt.Println("Переменная s имеет тип bool")
	case chan int:
		fmt.Println("Переменная s имеет тип chan int")
	default:
		fmt.Println("Неизвестный тип переменной s")
	}

	var b interface{} = true
	switch b.(type) {
	case int:
		fmt.Println("Переменная b имеет тип int")
	case string:
		fmt.Println("Переменная b имеет тип string")
	case bool:
		fmt.Println("Переменная b имеет тип bool")
	case chan int:
		fmt.Println("Переменная b имеет тип chan int")
	default:
		fmt.Println("Неизвестный тип переменной b")
	}
	// Для проверки канала используем пакет reflect и функцию kind
	var c interface{} = make(chan int)
	switch reflect.TypeOf(c).Kind() {
	case reflect.Chan:
		fmt.Println("Переменная c имеет тип chan int")
	default:
		fmt.Println("Неизвестный тип переменной c")
	}
}
