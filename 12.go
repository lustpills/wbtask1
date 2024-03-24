package main

import "fmt"

// Функция createSet создает собственное множество для данной последовательности строк.
func createSet(input []string) map[string]bool {
	set := make(map[string]bool) // Инициализируем пустое множество

	// Проходим по каждой строке во входной последовательности
	for _, str := range input {
		set[str] = true // Добавляем строку в множество
	}

	return set // Возвращаем собственное множество
}

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	set := createSet(sequence) // Создаем собственное множество для данной последовательности строк

	fmt.Println("Собственное множество:")
	for key := range set {
		fmt.Println(key) // Выводим элементы множества на экран
	}
}
