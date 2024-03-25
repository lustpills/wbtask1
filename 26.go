package main

import (
	"fmt"
	"strings"
)

// Функция, которая проверяет, что все символы в строке уникальные (регистронезависимо)
func areAllCharactersUnique(input string) bool {
	seen := make(map[rune]bool)

	// Приведем строку к нижнему регистру для регистронезависимой проверки
	input = strings.ToLower(input)

	for _, char := range input {
		if seen[char] {
			return false
		}
		seen[char] = true
	}

	return true
}

func main() {
	input := "abcd "

	result := areAllCharactersUnique(input)

	fmt.Printf("Все символы в строке уникальные: %v\n", result)
}
