package main

import (
	"fmt"
	"sync"
)

/*
Написать программу, которая конкурентно рассчитает
значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

func task2() {
	var a = [5]int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	for i := range a {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(a[i] * a[i])
		}()
		wg.Wait()
	}
}
