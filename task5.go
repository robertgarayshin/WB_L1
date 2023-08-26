package main

import (
	"fmt"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
*/
func task5() {
	ch := make(chan int64)
	n := 0
	fmt.Print("Enter working time (in seconds): ")
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}

	go func() {
		for {
			now := time.Now().Unix()
			fmt.Printf("Sent %d\n", now)
			ch <- now
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			fmt.Printf("Received %d\n", <-ch)
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(time.Duration(n) * time.Second)
}
