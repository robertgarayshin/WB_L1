package main

import (
	"context"
	"fmt"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.
func task6() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		// Способ 1: по завершении работы
		defer wg.Done()
		fmt.Println("Goroutine started")
		time.Sleep(1500 * time.Millisecond)
		fmt.Println("Goroutine stopped")
	}()
	wg.Wait()

	stop := make(chan bool)
	defer close(stop)
	go func(stop chan bool) {
		// Способ 2: через канал с сигналом Stop
		fmt.Println("Goroutine started")
		for {
			select {
			// Как только в канале что-то появилось, останавливаем горутину
			case <-stop:
				fmt.Println("Goroutine stopped")
				return
			default:
				fmt.Println(time.Now().Unix())
				time.Sleep(1000 * time.Millisecond)
			}
		}
	}(stop)
	time.Sleep(3 * time.Second)
	stop <- true

	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT)
	wg.Add(1)
	go func(ctx context.Context) {
		// Способ 3: через контекст
		fmt.Println("Goroutine started")
		for {
			select {
			case <-ctx.Done():
				// Если получаем SIGINT от ОС, то посылаем сигнал об остановке
				fmt.Println("Goroutine stopped")
				wg.Done()
				return
			default:
				fmt.Println("I'm working")
				time.Sleep(time.Second)
			}
		}

	}(ctx)
	wg.Wait()
}
