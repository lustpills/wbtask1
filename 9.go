package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// Увеличиваем счетчик на 3, чтобы дождаться завершения трех горутин
	wg.Add(3)
	numbersCh := make(chan int) // Канал для передачи чисел
	squaredCh := make(chan int) // Канал для передачи квадратов чисел

	// Горутина для отправки чисел в канал numbersCh
	go func(out chan<- int) {
		defer wg.Done()
		numbers := []int{1, 2, 3, 4, 5}
		for _, num := range numbers {
			out <- num
		}
		close(out) // Закрываем канал после отправки всех чисел
	}(numbersCh)

	// Горутина для вычисления квадратов чисел и отправки их в канал squaredCh
	go func(out chan<- int, in <-chan int) {
		defer wg.Done()
		for num := range in {
			out <- num * num
		}
		close(out) // Закрываем канал после вычисления всех квадратов
	}(squaredCh, numbersCh)

	// Горутина для вывода результатов из канала squaredCh
	go func(in <-chan int) {
		defer wg.Done()
		for squaredNum := range in {
			fmt.Println(squaredNum)
		}
	}(squaredCh)

	wg.Wait() // Ожидаем завершения всех горутин
}
