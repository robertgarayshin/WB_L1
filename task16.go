package main

import (
	"fmt"
)

func swap16(a *[]int, i int, j int) {
	temp := (*a)[i]
	(*a)[i] = (*a)[j]
	(*a)[j] = temp
}

func partition(a *[]int, start int, end int) int {
	elem := (*a)[end]

	pInd := start

	for i := start; i < end; i++ {
		if (*a)[i] <= elem {
			swap16(a, i, pInd)
			pInd += 1
		}
	}

	swap16(a, end, pInd)
	return pInd
}
func quicksort(a *[]int, start int, end int) {
	if start >= end {
		return
	}

	elem := partition(a, start, end)

	quicksort(a, start, elem-1)
	quicksort(a, elem+1, end)
}
func task16() {
	var a = []int{8, -1, 12, 86, 2, -6, 7, 0, -8, 1024}
	fmt.Printf("Before:\t\t")
	fmt.Println(a)
	quicksort(&a, 0, len(a)-1)
	fmt.Printf("After:\t\t")
	fmt.Println(a)
}
