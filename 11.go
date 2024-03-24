package main

import (
	"fmt"
)

// Функция intersection принимает два множества в виде map[int]bool и возвращает их пересечение.
func intersection(set1, set2 map[int]bool) map[int]bool {
	result := make(map[int]bool)

	// Перебираем элементы первого множества.
	for k := range set1 {
		// Если элемент присутствует во втором множестве, добавляем его в результат.
		if set2[k] {
			result[k] = true
		}
	}

	return result
}

func main() {
	// Создаем два множества с помощью map.
	var set1 map[int]bool //map[int]bool{1: true, 2: true, 3: true, 4: true}
	set2 := map[int]bool{3: true, 4: true, 5: true, 6: true}

	// Находим пересечение множеств и сохраняем результат в переменной intersect.
	intersect := intersection(set1, set2)

	fmt.Println("Пересечение множеств:")
	// Выводим результат пересечения на экран.
	for k := range intersect {
		fmt.Println(k)
	}
}
