package main

import (
	"log"
	"time"
)

// Реализовать собственную функцию sleep.
func sleep(d time.Duration) time.Time {
	// time.After() возвращает в канал текущее время по истечении промежутка в параметре.
	return <-time.After(d)
}

func sleep2(d time.Duration) time.Time {
	ticker := time.Tick(d)
	// Функция-тикер вернет канал с доставкой пакетов по истечении указанного времени (хотя вообще она будет "тикать"
	// каждый определенный промежуток времени). В нашем случае первый "возврат" тикера завершит функцию sleep2
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
