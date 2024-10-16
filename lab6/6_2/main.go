package main

import (
	"fmt"
)

func fibonacci(n int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c) // Закрытие канала после завершения генерации

		x, y := 0, 1
		for i := 0; i < n; i++ {
			c <- x //Отправка следующего числа Фибоначчи в канал `c`.
			x, y = y, x+y
		}
	}()
	return c
}

func main() {
	c := fibonacci(10) // Получение канала с числами Фибоначчи

	for n := range c {
		fmt.Println(n)
	}
}
