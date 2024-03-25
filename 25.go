package main

import (
	"fmt"
	"time"
)

// Функция sleep принимает количество секунд в качестве аргумента
func sleep(seconds int) {
	<-time.After(time.Duration(seconds) * time.Second)
}

func main() {
	fmt.Println("Начало выполнения программы")

	// Вызываем функцию sleep для задержки выполнения на 3 секунды
	sleep(3)

	fmt.Println("Программа завершена")
}
