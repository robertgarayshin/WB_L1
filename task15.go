package main

import (
	"log"
	"os"
)

func createHugeString(i int, f interface{}) string { return "str" }

// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
// Приведите корректный пример реализации.
var justString string

func someFunc() {
	v := createHugeString(1<<10, nil)
	justString = v[:100]
}

func task15() {
	someFunc()
}

/*
Проблемы:
1. Строка 8: работа с большой строкой (то есть создание ее в функции). Большие строки занимают
много времени работы программы и могут не поместиться в память.
2. Строка 9: попытка изменения строки приводит к созданию новой строки, что снова занимает время работы и память.
3. Строка 5: неоправданное использование глобальной переменной. Может привести к непредсказуемому изменению
значения переменной в процессе работы программы. Также, глобальные переменные создаются в куче, время обращения к
которой больше времени обращения к стеку, где создаются локальные переменные.
*/

func solution1() {
	f, err := os.Create("huge_string.txt")
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)
	_ = createHugeString(1<<30, f)
}

func solution2() {
	b := []byte(createHugeString(1<<30, nil))
	b = b[:100]
}

func solution3() string {
	return createHugeString(1<<10, nil)[:100]
}
