package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

func main() {

	// способ 1
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("Горутина 1 останавливается...")
				return
			default:
				// Выполняем полезную работу...
			}
		}
	}()

	time.Sleep(time.Second)
	stop <- true

	// способ 2

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Горутина 2 останавливается...")
				return
			default:
				// Выполняем полезную работу...
			}
		}
	}()

	time.Sleep(time.Second)
	cancel()

	// 3

	stop1 := make(chan struct{})

	go func() {
		for {
			select {
			case <-stop1:
				fmt.Println("Горутина 3 останавливается...")
				return
			default:
				// Выполняем полезную работу...
			}
		}
	}()

	time.Sleep(time.Second)
	close(stop1)

	// 4
	var running int32 = 1

	go func() {
		for atomic.LoadInt32(&running) == 1 {
			// Выполнение работы...
		}
		fmt.Println("stopped 4")
	}()

	// Для остановки
	atomic.StoreInt32(&running, 0)

	select {}
}
