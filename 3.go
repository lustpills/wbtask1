package main

import (
	"fmt"
	"sync"
)

func main() {
	// Массив чисел, для которых нужно вычислить квадраты
	numbers := []int{2, 4, 6, 8, 10}

	// Канал для передачи результатов вычислений
	results := make(chan int)

	// WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// Добавляем количество задач в WaitGroup равное количеству чисел
	wg.Add(len(numbers))

	// Запускаем горутину для каждого числа в массиве
	for _, number := range numbers {
		go func(n int) {
			defer wg.Done()   // Уменьшаем счетчик WaitGroup после завершения горутины
			square := n * n   // Вычисляем квадрат числа
			results <- square // Отправляем результат в канал
		}(number)
	}

	// Запускаем горутину для чтения результатов из канала
	go func() {
		wg.Wait()      // Ожидаем завершения всех горутин
		close(results) // Закрываем канал после получения всех результатов
	}()

	// Читаем из канала, складываем все в sum, выводим в консоль
	sum := 0
	for result := range results {
		sum += result
	}
	fmt.Println(sum)
}
