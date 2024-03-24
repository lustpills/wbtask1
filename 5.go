package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	duration := 5 * time.Second // N секунд, после которых программа должна завершиться
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel() // Освобождаем ресурсы контекста по завершению работы

	dataChannel := make(chan interface{})

	// Горутина для записи значений в канал
	go func(ctx context.Context) {
		counter := 0
		for {
			select {
			case <-ctx.Done():
				close(dataChannel) // Закрываем канал, чтобы сигнализировать о завершении
				return
			default:
				dataChannel <- counter
				counter++
				time.Sleep(1 * time.Second) // Имитация затрат времени на генерацию данных
			}
		}
	}(ctx)

	// Горутина для чтения значений из канала
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Время вышло, программа завершается...")
				return
			case data, ok := <-dataChannel:
				if !ok {
					fmt.Println("Канал закрыт, программа завершается...")
					return
				}
				fmt.Printf("Прочитано значение: %d\n", data)
			}
		}
	}(ctx)

	// Дожидаемся истечения времени контекста
	<-ctx.Done()
}
