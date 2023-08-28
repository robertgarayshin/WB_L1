package main

import (
	"fmt"
)

// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

// swap16 для обмена переменных местами
func swap16(a *[]int, i int, j int) {
	temp := (*a)[i]
	(*a)[i] = (*a)[j]
	(*a)[j] = temp
}

// Функция деления элементов слайса
func partition(a *[]int, start int, end int) int {
	// выбираем опорный элемент
	elem := (*a)[end]

	// выбираем индекс для смещения
	pInd := start

	for i := start; i < end; i++ {
		// если текущий элемент меньше или равен опорному
		if (*a)[i] <= elem {
			swap16(a, i, pInd) // сдвигаем его левее индекса для смещения
			pInd += 1
		}
	}
	// меняем местами конец и индекс для смещения
	swap16(a, end, pInd)
	return pInd
}
func quicksort(a *[]int, start int, end int) {
	if start >= end {
		return // условие завершения рекурсии
	}

	elem := partition(a, start, end) // деление среза

	quicksort(a, start, elem-1) // сортировка левой части разделенного слайса
	quicksort(a, elem+1, end)   // сортировка правой части разделенного слайса
}
func task16() {
	var a = []int{8, -1, 12, 86, 2, -6, 7, 0, -8, 1024}
	fmt.Printf("Before:\t\t")
	fmt.Println(a)
	quicksort(&a, 0, len(a)-1)
	fmt.Printf("After:\t\t")
	fmt.Println(a)
}
