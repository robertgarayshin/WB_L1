package main

import (
	"fmt"
	"unicode/utf8"
)

func reverse1(s string) string {
	a := []rune(s)
	var res []rune
	for i := len(a) - 1; i >= 0; i-- {
		res = append(res, a[i])
	}
	return string(res)
}

func reverse2(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func reverse3(s string) string {
	o := make([]rune, utf8.RuneCountInString(s))
	i := len(o)
	for _, c := range s {
		i--
		o[i] = c
	}
	return string(o)
}

func task19() {
	s := "главрыба"
	res := reverse1(s)
	fmt.Println(res)
	res = reverse2(s)
	fmt.Println(res)
	res = reverse3(s)
	fmt.Println(res)
}
