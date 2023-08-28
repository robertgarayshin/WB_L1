package main

import "fmt"

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
func task12() {
	var strs = []string{"cat", "cat", "dog", "cat", "tree"}
	// создаем структуру множества
	set := make(map[string]interface{})
	for i := range strs {
		set[strs[i]] = nil
	}

	// ключи - строки (множество), значения - пустые
	for key, _ := range set {
		fmt.Printf("%s ", key)
	}
}
