package main

import (
	"fmt"
	"sync"
)

/*
Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
*/

func task3() {
	var a = []int{2, 4, 6, 8, 10}
	sum := 0
	// Создаем канал для "общения между горутинами
	ch := make(chan int)
	var wg sync.WaitGroup

	for i := range a {
		wg.Add(1)
		go func() {
			// Записываем результат в канал
			ch <- a[i] * a[i]
		}()
		// Извлекаем результат из канала для подсчета суммы
		sum += <-ch
	}
	close(ch)
	fmt.Println(sum)
}
