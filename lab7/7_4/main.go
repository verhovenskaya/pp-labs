package main

//4.	Создание HTTP-сервера:
//•	Реализуйте базовый HTTP-сервер с обработкой простейших GET и POST запросов.
//•	Сервер должен поддерживать два пути:
//•	GET /hello — возвращает приветственное сообщение.
//•	POST /data — принимает данные в формате JSON и выводит их содержимое в консоль.

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/data", dataHandler)

	fmt.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Считываем тело запроса
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Десериализуем JSON-данные
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	// Выводим данные в консоль
	fmt.Println("Received data:", data)

	// Отправляем ответ клиенту
	fmt.Fprintln(w, "Data received successfully.")
}
