package main

//Применение select для управления каналами:
//•	Создайте две горутины, одна из которых будет генерировать случайные числа, а другая — отправлять сообщения об их чётности/нечётности.
//•	Используйте конструкцию select для приёма данных из обоих каналов и вывода результатов в консоль.
//•	Продемонстрируйте, как select управляет многоканальными операциями.

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomNumbers(ch chan int) {
	defer close(ch)
	for i := 0; i < 10; i++ {
		ch <- rand.Intn(100)
		time.Sleep(500 * time.Millisecond)
	}
}

func checkParity(number int, parityCh chan string) {

	if number%2 == 0 {
		parityCh <- "Чётное"
	} else {
		parityCh <- "Нечётное"
	}
}

func main() {
	randCh := make(chan int)
	parityCh := make(chan string)

	go generateRandomNumbers(randCh)

	for i := 0; i < 10; i++ {
		select {
		case number := <-randCh:
			fmt.Println("Получено число:", number)
			go checkParity(number, parityCh) // Проверяем четность в отдельной горутине
		case parity := <-parityCh:
			fmt.Println("Чётность:", parity)
		}
	}
}
