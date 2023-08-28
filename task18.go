package main

import (
	"fmt"
	"sync"
)

type counter struct {
	// структура-счетчик
	sync.Mutex
	cnt int
}

func (c *counter) Inc() {
	// инкрементация счетчика
	c.Lock()
	defer c.Unlock()

	c.cnt++
}

func (c *counter) getValue() int {
	// получение значения счетчика
	return c.cnt
}

// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.
func task18() {
	var c = &counter{cnt: 0} // инициализация
	stop := make(chan bool)  // канал сигналов завершения
	go work(c, stop)
	select {
	case <-stop: // если пришел сигнал о завершении работы
		fmt.Printf("%d workers finished their work", c.cnt)
	}

}

func work(cnt *counter, stop chan bool) {
	// запускаем работу
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(num int, cnt *counter, wg *sync.WaitGroup) {
			// добавляем горутину с некоторой задачей
			defer wg.Done()

			fmt.Printf("Worker %d starting\n", num)
			cnt.Inc() // увеличиваем значение счетчика
			fmt.Printf("Worker %d done\n", num)
		}(i, cnt, &wg)
	}
	wg.Wait()    // дожидаемся завершения работы
	stop <- true // посылаем сигнал о завершении
	close(stop)  // закрываем канал
}
