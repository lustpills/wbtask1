package main

import (
	"fmt"
	//"unicode/utf8"
)

// reverseString функция для переворачивания строки
func reverseString(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	input := "главрыба😀"
	reversed := reverseString(input)
	fmt.Println("Reversed string:", reversed)
}
