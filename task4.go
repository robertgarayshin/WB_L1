package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"
	"time"
)

/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/

func worker(n int, ch chan int64, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			data, ok := <-ch
			if !ok {
				log.Fatal("Channel closed")
				return
			}
			log.Println(data, "saved")
		default:
			log.Printf("Worker %d received data %d", n, <-ch)
		}
	}
}

func task4() {
	var workersCount int
	ch := make(chan int64)
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT)
	fmt.Print("Enter workers count: ")
	_, err := fmt.Scanf("%d", &workersCount)
	if err != nil {
		return
	}
	for i := 0; i < workersCount; i++ {
		go worker(i, ch, ctx)
	}
	for {
		select {
		case <-ctx.Done():
			cancel()
			log.Println("CANCEL")
			return
		default:
			ch <- time.Now().Unix()
			time.Sleep(time.Second)
		}
	}
}
