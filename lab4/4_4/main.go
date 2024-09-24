package main

//Написать программу, которая считывает несколько чисел, введенных пользователем, и выводит их сумму.
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите числа, разделенные пробелами: ")
	input, _ := reader.ReadString('\n')
	// Удаление символа новой строки из ввода
	input = strings.TrimSpace(input)

	// Разделение строки на отдельные числа
	numbers := strings.Split(input, " ")

	// Вычисление суммы
	var sum int
	for _, numberStr := range numbers {
		number, _ := strconv.Atoi(numberStr)
		sum += number
	}

	// Вывод суммы
	fmt.Println("Сумма чисел:", sum)
}
