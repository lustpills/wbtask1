package main

import (
	"fmt"
	"sync"
)

// SafeMap - потокобезопасная обертка вокруг map.
type SafeMap struct {
	mu   sync.Mutex
	data map[string]int
}

// NewSafeMap создает новый экземпляр SafeMap.
func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[string]int),
	}
}

// Set устанавливает значение по ключу с блокировкой для потокобезопасности.
func (sm *SafeMap) Set(key string, value int) {
	sm.mu.Lock()
	sm.data[key] = value
	sm.mu.Unlock()
}

// Get возвращает значение по ключу с блокировкой для потокобезопасности.
func (sm *SafeMap) Get(key string) (int, bool) {
	sm.mu.Lock()
	val, ok := sm.data[key]
	sm.mu.Unlock()
	return val, ok
}

func main() {
	sm := NewSafeMap()

	// Параллельное добавление данных в SafeMap из нескольких горутин.
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sm.Set(fmt.Sprintf("key%d", i), i)
		}(i)
	}
	wg.Wait()

	// Получение и печать значения по ключу после завершения всех горутин.
	if value, exists := sm.Get("key99"); exists {
		fmt.Println("Value:", value)
	}
}
