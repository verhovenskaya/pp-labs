package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

// Задача: реверсирование строки
type Task struct {
	Line string
}

// Результат: реверсированная строка
type Result struct {
	Line string
	Err  error
}

// Функция, выполняющая реверсирование строки
func reverseString(task Task) Result {
	reversed := strings.Split(task.Line, "")
	for i, j := 0, len(reversed)-1; i < j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = reversed[j], reversed[i]
	}
	return Result{Line: strings.Join(reversed, ""), Err: nil}
}

// Функция воркера
func worker(taskChan chan Task, resultChan chan Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range taskChan {
		result := reverseString(task)
		resultChan <- result
	}
}

func main() {
	// Получение имени файла и количества воркеров от пользователя
	fmt.Print("Введите имя файла: ")
	var filename string
	fmt.Scanln(&filename)

	fmt.Print("Введите количество воркеров: ")
	var numWorkers int
	fmt.Scanln(&numWorkers)

	// Открытие файла
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	// Каналы для задач и результатов
	taskChan := make(chan Task)
	resultChan := make(chan Result)

	// Создание пула воркеров
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker(taskChan, resultChan, &wg)
	}

	// Чтение строк из файла и отправка задач в канал
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		taskChan <- Task{Line: scanner.Text()}
	}
	// Закрываем канал задач после того, как прочитали все строки из файла
	close(taskChan)

	// Получение результатов из канала и вывод в консоль
	wg.Wait()
	for i := 0; i < numWorkers; i++ {
		result := <-resultChan
		if result.Err != nil {
			fmt.Println("Ошибка:", result.Err)
		} else {
			fmt.Println("Реверсированная строка:", result.Line)
		}
	}
}
