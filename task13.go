package main

import "fmt"

func swap(a int, b int) (int, int) {
	return b, a
}

func task13() {
	a := 1
	b := 2
	fmt.Printf("Before func swap a = %d, b = %d\n", a, b)
	a, b = swap(a, b)
	fmt.Printf("After func swap a = %d, b = %d\n", a, b)

	a = 1
	b = 2
	fmt.Printf("Before sum swap a = %d, b = %d\n", a, b)
	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("After sum swap a = %d, b = %d", a, b)
}
