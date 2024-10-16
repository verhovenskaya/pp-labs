//Разработка многопоточного калькулятора:
//•	Напишите многопоточный калькулятор, который одновременно может обрабатывать запросы на выполнение простых операций (+, -, *, /).
//•	Используйте каналы для отправки запросов и возврата результатов.
//•	Организуйте взаимодействие между клиентскими запросами и серверной частью калькулятора с помощью горутин.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Запрос на выполнение операции
type request struct {
	num1   float64
	num2   float64
	op     string
	result chan float64
}

// Обработчик запросов
func handleRequest(req request) {
	switch req.op {
	case "+":
		req.result <- req.num1 + req.num2
	case "-":
		req.result <- req.num1 - req.num2
	case "*":
		req.result <- req.num1 * req.num2
	case "/":
		if req.num2 == 0 {
			req.result <- 0 // Деление на ноль
		} else {
			req.result <- req.num1 / req.num2
		}
	default:
		req.result <- 0 // Неизвестная операция
	}
}

func main() {
	requests := make(chan request) // Канал для запросов

	// Запуск горутин для обработки запросов
	for i := 0; i < 4; i++ { // 4 горутины для обработки запросов
		go func() {
			for req := range requests {
				handleRequest(req)
			}
		}()
	}

	// Клиентская часть:
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите выражение (например, 10 + 5):")
	for scanner.Scan() {
		input := scanner.Text()
		parts := strings.Split(input, " ")
		if len(parts) != 3 {
			fmt.Println("Неверный формат выражения. Введите 2 числа и операцию.")
			continue
		}

		num1, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			fmt.Println("Ошибка преобразования первого числа.")
			continue
		}
		num2, err := strconv.ParseFloat(parts[2], 64)
		if err != nil {
			fmt.Println("Ошибка преобразования второго числа.")
			continue
		}

		result := make(chan float64)
		req := request{
			num1:   num1,
			num2:   num2,
			op:     parts[1],
			result: result,
		}

		requests <- req

		go func() {
			res := <-result
			fmt.Println("Результат:", res)
		}()
	}

	close(requests) // Закрытие канала запросов
}
