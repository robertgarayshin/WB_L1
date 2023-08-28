package main

import (
	"fmt"
)

// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
// во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
func writer(ch chan int, a [10]int) {
	defer close(ch)
	for i := range a {
		ch <- a[i] // записываем числа из массива
	}
}

func multiplier(x int, ch2 chan int) {
	ch2 <- x * 2 // записываем в другой канал результат умножения числа на 2
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
		x, ok := <-ch // получаем данные из канала
		if !ok {
			break // если канал закрыт - останавливаем выполнение
		} else {
			go multiplier(x, ch2) // если все ок - умножаем и выводим числа
			fmt.Println(<-ch2)
		}
	}

}
