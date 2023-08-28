package main

import (
	"fmt"
	"strings"
)

func rev(s *[]string) {
	for i, j := 0, len(*s)-1; i < len(*s)/2; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

func task20() {
	s := "snow dog sun"
	spl := strings.Split(s, " ")
	rev(&spl)
	fmt.Println(strings.Join(spl, " "))
}
