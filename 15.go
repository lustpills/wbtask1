package main

import "fmt"

func createHugeString(size int) string {
	// Логика создания огромной строки
	hugeString := "Huge String Content..."
	return hugeString
}

// исправление заключается в том,
// что огромная строка не хранится в глобальной переменной
func someFunc() string {
	v := createHugeString(1 << 10)
	return v[:100]
}

func main() {
	justString := someFunc()
	fmt.Println(justString)
}
