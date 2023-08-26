package main

import "fmt"

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
func task12() {
	var strs = []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]interface{})
	for i := range strs {
		set[strs[i]] = nil
	}

	for key, _ := range set {
		fmt.Printf("%s ", key)
	}
}
