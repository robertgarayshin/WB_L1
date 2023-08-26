package main

import (
	"fmt"
	"sync"
	"time"
)

//Реализовать конкурентную запись данных в map.

type MyMap struct {
	sync.Mutex
	m map[string]int
}

func NewMyMap() *MyMap {
	return &MyMap{
		m: make(map[string]int),
	}
}

func (a *MyMap) SetData(key string, val int) {
	a.Lock()
	defer a.Unlock()
	a.m[key] = val
}

func (a *MyMap) GetData(key string) (int, bool) {
	a.Lock()
	defer a.Unlock()
	val, ok := a.m[key]
	return val, ok
}

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
