package main

import (
	"fmt"
	"os"
)

func main() {
	// Создание карты
	people := make(map[string]int)

	// Добавление людей
	people["Иван"] = 30
	people["Мария"] = 25
	people["Петр"] = 40

	// Добавление нового человека
	var newName string
	var newAge int
	fmt.Println("Введите нового человека:")
	fmt.Fscan(os.Stdin, &newName)
	fmt.Println("Введите возраст:")
	fmt.Fscan(os.Stdin, &newAge)

	// Добавление нового человека в карту
	people[newName] = newAge

	// Вывод всех записей
	fmt.Println("Все люди:")
	for name, age := range people {
		fmt.Printf("%s - %d лет\n", name, age)
	}

	// Вычисление среднего возраста
	averageAge := calculateAverageAge(people)
	fmt.Printf("Средний возраст: %.2f\n", averageAge)
}

// Функция для вычисления среднего возраста
func calculateAverageAge(people map[string]int) float64 {
	var totalAge int
	var count int
	for _, age := range people {
		totalAge += age
		count++
	}
	if count == 0 {
		return 0.0
	}
	return float64(totalAge) / float64(count)
}
