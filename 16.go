package main

import (
	"fmt"
	"math/rand"
)

// Функция partition разделяет массив на три части: элементы меньше, равные и больше опорного элемента
func partition(arr []int, l, r, c int) (int, int) {
	e := l // Индекс конца части с элементами меньше опорного
	g := r // Индекс начала части с элементами больше опорного

	for i := l; i < g; {
		if arr[i] < c {
			arr[e], arr[i] = arr[i], arr[e]
			e++
			i++
		} else if arr[i] > c {
			g--
			arr[g], arr[i] = arr[i], arr[g]
		} else {
			i++
		}
	}

	return e, g
}

// Функция getRandomPivot возвращает случайный элемент из диапазона [l, r)
func getRandomPivot(l, r int) int {
	return rand.Intn(r-l) + l
}

// Функция quicksort выполняет сортировку массива методом быстрой сортировки
func quicksort(arr []int, l, r int) {
	if r-l >= 2 {
		pivotIndex := getRandomPivot(l, r)
		c := arr[pivotIndex] // Опорный элемент - случайный элемент
		a, b := partition(arr, l, r, c)
		quicksort(arr, l, a) // Сортировка левой части массива
		quicksort(arr, b, r) // Сортировка правой части массива
	}
}

func main() {
	arr := []int{3, 5, 2, 8, 6, 4, 5, 7, 11, 1, 1, 1, 13213213, 2, 4, 5}
	l := 0
	r := len(arr)

	quicksort(arr, l, r)

	fmt.Println(arr)
}
