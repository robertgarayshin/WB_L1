package main

import "fmt"

func findKeys(a []float32) []int {
	var result []int
	for i := range a {
		result = append(result, (int(a[i])/10)*10)
	}
	return result
}

//Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
//Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
func task10() {
	var a = []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	keys := findKeys(a)
	temp := make(map[int][]float32)
	for i := range keys {
		temp[keys[i]] = append(temp[keys[i]], a[i])
	}

	fmt.Println(temp)
}
