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
		defer wg.Done()
		fmt.Println("Goroutine started")
		time.Sleep(1500 * time.Millisecond)
		fmt.Println("Goroutine stopped")
	}()
	wg.Wait()

	stop := make(chan bool)
	defer close(stop)
	go func(stop chan bool) {
		fmt.Println("Goroutine started")
		for {
			select {
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
		fmt.Println("Goroutine started")
		for {
			select {
			case <-ctx.Done():
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
