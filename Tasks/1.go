package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (h *Human) SayHello() {
	fmt.Println("Hi, my name is", h.name)
}

type Action struct {
	*Human
}

func main() {
	h := &Human{name: "Slim Shady", age: 30}
	a := &Action{h}

	a.SayHello() // вызов метода SayHello у структуры Human через структуру Action
}
