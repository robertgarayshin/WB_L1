package main

import (
	"fmt"
	"strings"
)

func check(s string) bool {
	s = strings.ToLower(s)
	set := make(map[rune]interface{})
	for i := 0; i < len(s); i++ {
		set[rune(s[i])] = ""
	}
	if len(set) < len(s) {
		return false
	}
	return true
}

func task26() {
	fmt.Println(check(""))
}
