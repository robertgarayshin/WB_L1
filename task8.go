package main

import (
	"fmt"
	"strconv"
)

func variant1(a int64, bitPos int) {
	b := fmt.Sprintf("%b", a)
	fmt.Println(b)
	newStr := b[:len(b)-bitPos-1] + func() string {
		if b[len(b)-bitPos-1] == '1' {
			return "0"
		} else {
			return "1"
		}
	}() + b[len(b)-bitPos:]
	newInt, _ := strconv.ParseInt(newStr, 2, 64)
	fmt.Printf("variant1: %d\n", newInt)
}

func variant2(a int64, bitPos int) {
	var mask int64
	mask = 1 << bitPos
	if a|mask == a {
		fmt.Printf("variant2: %d", a&^mask)
	} else {
		fmt.Printf("variant2: %d", a|mask)
	}

}

// Дана переменная int64. Разработать программу, которая устанавливает i-й бит в 1 или 0.
func task8() {
	var a int64
	var bitPos int
	_, err := fmt.Scan(&a, &bitPos)
	if err != nil {
		return
	}
	variant1(a, bitPos)
	variant2(a, bitPos)
}
