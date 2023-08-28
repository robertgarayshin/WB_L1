package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

func rev(s *[]string) {
	// если можно инвертировать буквы в строке, то можно и строки в срезе строк с использованием обмена значений
	for i, j := 0, len(*s)-1; i < len(*s)/2; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

func task20() {
	s := "snow dog sun"
	spl := strings.Split(s, " ")        // делим строку на срез
	rev(&spl)                           // инвертируем порядок слов
	fmt.Println(strings.Join(spl, " ")) // обратно создаем единую строку из среза
}
