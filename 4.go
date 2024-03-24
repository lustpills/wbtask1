package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"

	//"strconv"
	"syscall"
	"time"
)

func main() {
	var numOfWorkers int

	fmt.Scanf("%d", &numOfWorkers)

	ctx, cancel := context.WithCancel(context.Background())

	// Подготовка к перехвату сигнала SIGINT (Ctrl+C)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)

	dataChannel := make(chan interface{})
	var wg sync.WaitGroup // Используется для ожидания завершения всех воркеров

	go func() {
		<-sigChan // Ожидание сигнала
		fmt.Println("\nReceived an interrupt signal, shutting down...")
		close(dataChannel) // Закрытие канала для остановки воркеров
		cancel()           // Отмена контекста для корректного завершения всех горутин
	}()

	go func(ctx context.Context) {
		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				return
			case dataChannel <- i:
				time.Sleep(1 * time.Second)
			}
		}
	}(ctx)

	for i := 0; i < numOfWorkers; i++ {
		wg.Add(1) // Увеличиваем счетчик группы ожидания
		go func(id int, ctx context.Context, dataChan <-chan interface{}) {
			defer wg.Done() // Уменьшаем счетчик группы ожидания при выходе из горутины
			for {
				select {
				case <-ctx.Done():
					fmt.Printf("Worker %d shutting down...\n", id)
					return
				case data, ok := <-dataChan:
					if !ok {
						fmt.Printf("Worker %d shutting down...\n", id)
						return
					}
					fmt.Printf("Worker %d received data: %d\n", id, data)
				}
			}
		}(i+1, ctx, dataChannel)
	}

	wg.Wait() // Ожидание завершения всех воркеров
	fmt.Println("All workers have been shut down. Exiting.")
}
