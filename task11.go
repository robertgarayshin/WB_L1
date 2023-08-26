package main

import "fmt"

type Intersection map[int]int

func (i *Intersection) Intersect(set1 []int, set2 []int) {
	for k := range set1 {
		(*i)[set1[k]] = 0
	}
	for k := range set2 {
		_, found := (*i)[set2[k]]
		if found {
			(*i)[set2[k]] += 1
		}
	}

	for key, value := range *i {
		if value < 1 {
			delete(*i, key)
		}
	}
}

// Реализовать пересечение двух неупорядоченных множеств.
func task11() {
	var set1 = []int{1, 2, 5, 7, 12, 10, 5, 1, 3, 0, -2}
	var set2 = []int{2, 8, 1, 3, -2, 6, 30, 34}
	var inters = make(Intersection)
	inters.Intersect(set1, set2)
	for key, _ := range inters {
		fmt.Printf("%d ", key)
	}

}
