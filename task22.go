package main

import (
	"fmt"
	"math"
	"math/big"
)

func variant1_22() {
	// Если числа могут поместиться в Int64, можно попробовать использовать их
	// Самый ненадежный вариант, так как можно получить переполнение в любой момент
	fmt.Println(math.MaxInt64)

	fmt.Printf("Minimum value of numbers %f", math.Pow(2, 20))
}

func variant2_22() {
	// Использование математики float
	// Надежен при работе с экспоненциальной формой числа
	// Но стоит вывести число на экран в привычном формате - сразу будет заметна недостаточная точность float
	a := 24e18
	b := 25e20
	fmt.Println("Variant with exponents")
	fmt.Println(a + b)
	fmt.Println(b - a)
	fmt.Println(a * b)
	fmt.Println(b / a)
	fmt.Println()
}

func variant3_22() {
	// Используем пакет Big
	// Самый медленный и громоздкий вариант
	fmt.Println("Variant with Big package")
	a := new(big.Int)
	a.SetString("24000000000000000000"+
		"", 10)
	b := new(big.Int)
	b.SetString("2500000000000000000000", 10)

	resAdd := new(big.Int)
	resAdd.Add(a, b)
	fmt.Println(resAdd)
	resSub := new(big.Int)
	resSub.Sub(b, a)
	fmt.Println(resSub)
	resMul := new(big.Int)
	resMul.Mul(a, b)
	fmt.Println(resMul)

	// Деление с использованием BigInt только целочисленное
	resDiv := new(big.Int)
	resDiv.Div(b, a)
	fmt.Println(resDiv)

}

func variant4_22() {
	// Не подойдет для вывода, но будет работать
	const (
		a   = 24000000000000000000
		b   = 2500000000000000000000
		add = a + b
		sub = b - a
		mul = a * b
		div = b / a
	)
}

func task22() {
	variant2_22()
	variant3_22()

	variant4_22()
}
