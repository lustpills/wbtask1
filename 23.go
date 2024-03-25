package main

import "fmt"

func main() {
	// Создаем слайс с некоторыми элементами
	slice := []int{1, 2, 3, 4, 5}

	// Выводим исходный слайс
	fmt.Println("Исходный слайс:", slice)

	// Удаляем i-ый элемент из слайса (например, 2-ой элемент)
	indexToRemove := 2
	slice = append(slice[:indexToRemove], slice[indexToRemove+1:]...)

	// Выводим слайс после удаления элемента
	fmt.Println("Слайс после удаления", indexToRemove, "-го элемента:", slice)
}
