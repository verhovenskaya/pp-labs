package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	port := "8080"

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Ошибка при создании слушателя:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Сервер запущен на порту", port)

	//канал для сигнала прерывания
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	// Цикл обработки подключений
	go func() {
		for {
			// Принимаем входящее соединение
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("Ошибка при приеме соединения:", err)
				continue
			}

			go handleConnection(conn)
		}
	}()

	// Ожидаем сигнала прерывания
	<-interrupt
	fmt.Println("Завершение работы сервера...")

	listener.Close()

	// Ожидаем завершения всех активных соединений
	fmt.Println("Ожидание завершения активных соединений...")
}

// Функция обработки соединения
func handleConnection(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {

		message := scanner.Text()
		fmt.Println("Получено сообщение от клиента:", message)

		fmt.Fprintf(conn, "Сообщение получено: %s\n", message)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка чтения данных от клиента:", err)
	}
}
