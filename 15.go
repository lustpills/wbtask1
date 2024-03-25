package main

import (
	"fmt"
	"strings"
)

func createHugeString(size int) string {
	str := strings.Repeat("a", size)
	return str
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
