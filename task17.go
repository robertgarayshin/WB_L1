package main

import "fmt"

func binarySearch(a *[]int, need int) int {
	left := 0
	right := len(*a) - 1
	res := -1

	for left <= right {
		mid := (left + right) / 2
		if need == (*a)[mid] {
			res = mid
			break
		}
		if need < (*a)[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return res
}

func task17() {
	var a = []int{8, -1, 12, 86, 2, -6, 7, 0, -8, 1024}
	var need int
	quicksort(&a, 0, len(a)-1)
	fmt.Print("Enter needed value: ")
	_, err := fmt.Scan(&need)
	if err != nil {
		return
	}
	fmt.Printf("Needed value index: %d", binarySearch(&a, need))
}
