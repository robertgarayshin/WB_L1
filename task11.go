package main

import "fmt"

// Intersection структура множества-пересечения
type Intersection map[int]uint8

func (i *Intersection) Intersect(set1 []int, set2 []int) {
	for k := range set1 { // берем все ключи из первого множества
		(*i)[set1[k]] = 0 // и устанавливаем их со значением 0
	}
	for k := range set2 {
		_, found := (*i)[set2[k]] // проходимся по второму множеству и проверяем
		if found {                // при нахождении уже существующего ключа
			(*i)[set2[k]] += 1 // добавляем значение
		}
	}

	for key, value := range *i {
		if value < 1 { // проверяем ключи на наличие значения меньше единицы (то есть число встретилось только в одном мн-ве)
			delete(*i, key) // если находим такое - удаляем
		}
	}
}

// Реализовать пересечение двух неупорядоченных множеств.
func task11() {
	var set1 = []int{1, 2, 5, 7, 12, 10, 5, 1, 3, 0, -2} // первое множество
	var set2 = []int{2, 8, 1, 3, -2, 6, 30, 34}          // второе множество
	var inters = make(Intersection)                      // пересечение
	inters.Intersect(set1, set2)                         // нахождение пересечения
	for key, _ := range inters {
		fmt.Printf("%d ", key)
	}

}
