package main

import (
	"fmt"
)

func main() {
	// Создаем слайс с температурами
	temperatures := []float64{-25.4, -27.0, -21.0, 13.0, 19.0, 15.5, 24.5, 32.5}

	// Создаем карту для группировки температур по диапазонам
	groups := make(map[int][]float64)

	// Итерируем по всем температурам и добавляем их в соответствующие группы
	for _, temp := range temperatures {
		group := int(temp) / 10 * 10                // Определяем группу для текущей температуры
		groups[group] = append(groups[group], temp) // Добавляем температуру в соответствующую группу
	}

	// Выводим группы температур
	for group, temps := range groups {
		fmt.Printf("%d: %v\n", group, temps)
	}
}
