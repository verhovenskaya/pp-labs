package main

//Написать программу, которая удаляет запись из карты по заданному имени.

import "fmt"

func main() {
	// Создание карты
	people := make(map[string]int)
	people["Иван"] = 30
	people["Мария"] = 25
	people["Петр"] = 40

	// Удаление записи по имени
	delete(people, "Иван")

	// Вывод оставшихся записей
	fmt.Println("Оставшиеся люди:")
	for name, age := range people {
		fmt.Printf("%s - %d лет\n", name, age)
	}
}
