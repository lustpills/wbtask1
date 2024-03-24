package main

import "fmt"

func main() {
	var a, b int

	// Считываем два целых числа с консоли
	fmt.Scanf("%d %d", &a, &b)

	// Обмен значениями переменных a и b без использования дополнительной переменной
	a = a + b
	b = -(b - a)
	a = a - b

	// Выводим значения переменных a и b
	fmt.Printf("\n%d %d\n", a, b)
}
