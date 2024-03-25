package main

import (
	"fmt"
)

// Рекурсивная функция для выполнения бинарного поиска в отсортированном массиве
func binarySearchRecursive(arr []int, target int, left, right int) int {
	// Базовый случай: если левая граница превышает правую границу, элемент не найден
	if left > right {
		return -1
	}

	mid := left + (right-left)/2

	// Если средний элемент равен целевому, возвращаем его индекс
	if arr[mid] == target {
		return mid
	} else if arr[mid] < target {
		// Если средний элемент меньше целевого, делаем поиск в правой половине массива
		return binarySearchRecursive(arr, target, mid+1, right)
	} else {
		// Иначе делаем поиск в левой половине массива
		return binarySearchRecursive(arr, target, left, mid-1)
	}
}

func main() {
	arr := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	target := 0
	index := binarySearchRecursive(arr, target, 0, len(arr)-1)
	if index != -1 {
		fmt.Printf("Element %d found at index %d\n", target, index)
	} else {
		fmt.Printf("Element %d not found in the array\n", target)
	}
}
