package main

import (
	"fmt"
	"strings"
)

func reverseWords(input string) string {
	words := strings.Fields(input) // Разбиваем строку на слова
	reversedWords := make([]string, 0, len(words))

	for i := len(words) - 1; i >= 0; i-- {
		reversedWords = append(reversedWords, words[i]) // Добавляем слова в обратном порядке
	}

	return strings.Join(reversedWords, " ") // Объединяем слова обратно в строку
}

func main() {
	input := "snow dog sun"
	reversed := reverseWords(input)
	fmt.Println("Reversed words:", reversed)
}
