package main

import (
	"fmt"
	"sync"
)

type counter struct {
	sync.Mutex
	cnt int
}

func (c *counter) Inc() {
	c.Lock()
	defer c.Unlock()

	c.cnt++
}

func (c *counter) getValue() int {
	return c.cnt
}

// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.
func task18() {
	var c = &counter{cnt: 0}
	stop := make(chan bool)
	go work(c, stop)
	select {
	case <-stop:
		fmt.Printf("%d workers finished their work", c.cnt)
	}

}

func work(cnt *counter, stop chan bool) {
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(num int, cnt *counter, wg *sync.WaitGroup) {
			defer wg.Done()

			fmt.Printf("Worker %d starting\n", num)
			cnt.Inc()
			fmt.Printf("Worker %d done\n", num)
		}(i, cnt, &wg)
	}
	wg.Wait()
	stop <- true
	close(stop)
}
