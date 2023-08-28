package main

import (
	"fmt"
	"strconv"
)

func variant1(a int64, bitPos int) {
	// Первый вариант - самый простой. Меняем ноль на единицу и наоборот в строковом представлении числа
	b := fmt.Sprintf("%b", a) // Сохраняем бинарное представление числа в виде строки
	fmt.Println(b)
	newStr := b[:len(b)-bitPos-1] + func() string {
		if b[len(b)-bitPos-1] == '1' {
			return "0"
		} else {
			return "1"
		}
	}() + b[len(b)-bitPos:] // меняем число в зависимости от значения бита в требуемой позиции
	newInt, _ := strconv.ParseInt(newStr, 2, 64) // сохраняем новое число
	fmt.Printf("variant1: %d\n", newInt)
}

func variant2(a int64, bitPos int) {
	// Второй вариант использует битовые операции и битовую маску.
	var mask int64
	// создаем маску на число в виде 0..00100..0
	mask = 1 << bitPos
	if a|mask == a { // проверяем, останется ли число таким же при использовании побитового ИЛИ
		fmt.Printf("variant2: %d", a&^mask) // если да, то используем инверсию бит
	} else {
		fmt.Printf("variant2: %d", a|mask) // иначе используем побитовое ИЛИ
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
