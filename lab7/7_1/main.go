// server.go
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	port := "8080"

	//слушатель для TCP-соединений
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Ошибка при создании слушателя:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Сервер запущен на порту", port)

	//обработка подключений
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Ошибка при приеме соединения:", err)
			continue
		}
		go handleConnection(conn)
	}
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
