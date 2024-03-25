package main

import (
	"fmt"
	"math"
)

// Структура Point для представления точки с координатами x и y
type Point struct {
	x, y float64
}

// Конструктор для создания новой точки
func NewPoint(x, y float64) Point {
	return Point{x, y}
}

// Метод для вычисления расстояния между двумя точками
func (p1 Point) Distance(p2 Point) float64 {
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
}

func main() {
	// Создаем две точки
	point1 := NewPoint(1, 2)
	point2 := NewPoint(4, 6)

	// Вычисляем расстояние между точками
	distance := point1.Distance(point2)

	// Выводим результат
	fmt.Printf("Расстояние между точкой (%.2f, %.2f) и точкой (%.2f, %.2f) составляет %.2f\n", point1.x, point1.y, point2.x, point2.y, distance)
}
