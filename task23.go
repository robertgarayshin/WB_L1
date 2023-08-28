package main

import "fmt"

func variant1_23() []string {
	a := []string{"A", "B", "C", "D", "E"}
	i := 2

	// Удалить элемент по индексу i из a.

	// 1. Копировать последний элемент в индекс i.
	a[i] = a[len(a)-1]

	// 2. Удалить последний элемент (записать нулевое значение).
	a[len(a)-1] = ""

	// 3. Усечь срез.
	a = a[:len(a)-1]

	return a // [A B E D]
}

func variant2_23() []string {
	a := []string{"A", "B", "C", "D", "E"}
	i := 2

	// Удалить элемент по индексу i из a.

	// 1. Выполнить сдвиг a[i+1:] влево на один индекс.
	copy(a[i:], a[i+1:])

	// 2. Удалить последний элемент (записать нулевое значение).
	a[len(a)-1] = ""

	// 3. Усечь срез.
	a = a[:len(a)-1]

	return a // [A B D E]
}

func task23() {
	fmt.Println(variant1_23())
	fmt.Println(variant2_23())
}
