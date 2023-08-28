package main

import "fmt"

// Реализовать бинарный поиск встроенными методами языка.
func binarySearch(a *[]int, need int) int {
	left := 0            // левая граница
	right := len(*a) - 1 // правая граница
	res := -1            // индекс искомого

	for left <= right { // если левая граница перешла правую - возвращаем результат
		mid := (left + right) / 2 // делим массив пополам
		if need == (*a)[mid] {
			res = mid // если середина - искомое, возвращаем
			break
		}
		if need < (*a)[mid] {
			right = mid - 1 // если искомое меньше середины, то смещаем правую границу ниже середины
		} else {
			left = mid + 1 // иначе смещаем левую границу выше середины
		}
	}
	return res
}

func task17() {
	var a = []int{8, -1, 12, 86, 2, -6, 7, 0, -8, 1024} // неупорядоченный массив a
	var need int                                        // искомое значение
	quicksort(&a, 0, len(a)-1)                          // используем быструю сортировку из прошлого задания
	fmt.Print("Enter needed value: ")
	_, err := fmt.Scan(&need)
	if err != nil {
		return
	}
	fmt.Printf("Needed value index: %d", binarySearch(&a, need)) // ищем индекс искомого
}
