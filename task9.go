package main

import (
	"fmt"
)

// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
// во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
func writer(ch chan int, a [10]int) {
	defer close(ch)
	for i := range a {
		ch <- a[i]
	}
}

func multiplier(x int, ch2 chan int) {
	ch2 <- x * 2
}

func task9() {
	var a [10]int
	for i := 0; i < 10; i++ {
		a[i] = i
	}
	ch := make(chan int)
	ch2 := make(chan int)
	go writer(ch, a)
	for {
		x, ok := <-ch
		if !ok {
			break
		} else {
			go multiplier(x, ch2)
			fmt.Println(<-ch2)
		}
	}

}
