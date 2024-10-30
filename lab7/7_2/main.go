// client.go
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	address := "localhost"
	port := "8080"

	conn, err := net.Dial("tcp", address+":"+port)
	if err != nil {
		fmt.Println("Ошибка при подключении к серверу:", err)
		os.Exit(1)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	// Цикл ввода и отправки сообщений
	for {

		fmt.Println("Введите сообщение (или 'exit' для выхода):")
		message, _ := reader.ReadString('\n')
		message = strings.TrimSpace(message)

		if message == "exit" {
			break
		}

		fmt.Fprintf(conn, "%s\n", message)

		scanner := bufio.NewScanner(conn)
		if scanner.Scan() {
			response := scanner.Text()
			fmt.Println("Ответ от сервера:", response)
		} else {
			fmt.Println("Ошибка при чтении ответа от сервера")
			break
		}
	}
}
