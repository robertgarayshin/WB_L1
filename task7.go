package main

import (
	"fmt"
	"sync"
	"time"
)

//Реализовать конкурентную запись данных в map.

// MyMap Создаем структуру данных с мьютексом на бору
type MyMap struct {
	sync.Mutex
	m map[string]int
}

// NewMyMap Конструктор карты
func NewMyMap() *MyMap {
	return &MyMap{
		m: make(map[string]int),
	}
}

// SetData Сеттер данных
func (a *MyMap) SetData(key string, val int) {
	a.Lock()
	defer a.Unlock()
	a.m[key] = val
}

// GetData геттер данных
func (a *MyMap) GetData(key string) (int, bool) {
	a.Lock()
	defer a.Unlock()
	val, ok := a.m[key]
	return val, ok
}

// Суть работы с Map в конкурентной среде заключается в использовании мьютекса для регулировки доступа к общему ресурсу
func task7() {
	a := NewMyMap()

	go func() {
		key := "ao"
		val := 1
		a.SetData(key, val)
	}()

	go func() {
		key := "bo"
		val := 2
		a.SetData(key, val)
	}()

	go func() {
		key := "co"
		val := 3
		a.SetData(key, val)
	}()

	time.Sleep(time.Second)
	fmt.Println(a.GetData("co"))
}
