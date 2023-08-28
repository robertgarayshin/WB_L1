package main

import "fmt"

// Разработать программу, которая в рантайме способна определить тип
// переменной: int, string, bool, channel из переменной типа interface{}.
func task14() {
	// Для определения типа переменной в fmt.Printf можно использовать плейсхолдер %T, что покажет формат переменной
	var i interface{}
	i = 1
	fmt.Printf("%T\n", i)
	i = "ab"
	fmt.Printf("%T\n", i)
	i = false
	fmt.Printf("%T\n", i)
	i = make(chan int)
	fmt.Printf("%T", i)
}
