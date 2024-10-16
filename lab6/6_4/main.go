package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter int
var mutex sync.Mutex // Мьютекс для синхронизации доступа к счетчику

func incrementCounter(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 1000; i++ {
		// Блокировка мьютекса
		mutex.Lock()
		counter++
		// Разблокировка мьютекса
		mutex.Unlock()
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // Использование всех ядер процессора

	var wg sync.WaitGroup
	wg.Add(10) // Добавление 10 горутин

	// Запуск 10 горутин, увеличивающих счетчик с мьютексом
	for i := 0; i < 10; i++ {
		go incrementCounter(&wg)
	}

	wg.Wait() // Ожидание завершения всех горутин
	fmt.Println("Счетчик (с мьютексом):", counter)

	// Сброс счетчика
	counter = 0

	// Запуск 10 горутин, увеличивающих счетчик без мьютекса
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter++
			}
		}()
	}

	wg.Wait() // Ожидание завершения всех горутин
	fmt.Println("Счетчик (без мьютекса):", counter)
}
