package main

import "fmt"

/*Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/

// Human struct declaration
type Human struct {
	Name string
	Age  int
}

// Human's method
func (h *Human) sayHi() {
	fmt.Printf("Hi! My name is %s", h.Name)
}

// Action struct declaration
type Action struct {
	ActionType string
	Human
}

func task1() {
	// Creating instance of structure Action
	var a = Action{
		ActionType: "saying",
		Human: Human{
			Name: "Slim Shady",
			Age:  25,
		},
	}
	// Using Human's method via Action structure instance
	a.sayHi()
}
