package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Функция для расчета факториала
func factorial(n int, wg *sync.WaitGroup) {
	defer wg.Done()             //уменьшает внутренний счетчик на единицу
	time.Sleep(2 * time.Second) // Имитация задержки
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	fmt.Printf("Факториал %d! = %d\n", n, result)
}

// Функция для генерации случайных чисел
func randomNumbers(count int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second) // Имитация задержки
	numbers := make([]int, count)
	for i := 0; i < count; i++ {
		numbers[i] = rand.Intn(100) // Генерация случайного числа от 0 до 99
	}
	fmt.Printf("Случайные числа: %v\n", numbers)
}

// Функция для вычисления суммы числового ряда
func sumSeries(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(3 * time.Second) // Имитация задержки
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Printf("Сумма ряда от 1 до %d = %d\n", n, sum)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(3) // Добавляем три горутины

	go factorial(5, &wg)     // Вычисляем факториал 5!
	go randomNumbers(5, &wg) // Генерируем 5 случайных чисел
	go sumSeries(10, &wg)    // Вычисляем сумму от 1 до 10

	wg.Wait() // Ожидаем завершения всех горутин
	fmt.Println("Все горутины завершены.")
}
