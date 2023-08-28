package main

import (
	"log"
	"time"
)

func sleep(d time.Duration) time.Time {
	return <-time.After(d)
}

func sleep2(d time.Duration) time.Time {
	ticker := time.Tick(d)
	for done := range ticker {
		return done
	}
	return time.Now()
}

func task25() {
	log.Println("start")
	sleep(2 * time.Second)
	log.Println("first sleep ends, starting second sleep")
	sleep2(2 * time.Second)
	log.Println("second sleep ends, stopping")
}
