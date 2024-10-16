package main

//1.	Создание TCP-сервера:
//•	Реализуйте простой TCP-сервер, который слушает указанный порт и принимает входящие соединения.
//•	Сервер должен считывать сообщения от клиента и выводить их на экран.
//•	По завершении работы клиенту отправляется ответ с подтверждением получения сообщения.

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var wg sync.WaitGroup

func handleConnection(conn net.Conn) {
	defer wg.Done()
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Ошибка чтения:", err)
			return
		}

		message := string(buffer[:n])
		fmt.Println("Получено сообщение:", message)

		_, err = conn.Write([]byte("Сообщение полученоn"))
		if err != nil {
			fmt.Println("Ошибка отправки:", err)
			return
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Сервер запущен на порту 8080")

	// Обработка завершения работы сервера
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-stopChan
		fmt.Println("nЗавершение работы сервера...")
		listener.Close()
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Ошибка при принятии соединения:", err)
			continue
		}

		wg.Add(1)
		go handleConnection(conn)
	}

	wg.Wait()
	fmt.Println("Все соединения закрыты. Сервер завершил работу.")
}
