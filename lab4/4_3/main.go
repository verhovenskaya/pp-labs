package main

//Написать программу, которая считывает строку с ввода и выводит её в верхнем регистре
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите строку: ")
	input, _ := reader.ReadString('\n')
	// Удаление символа новой строки из ввода
	input = strings.TrimSpace(input)

	// Вывод строки в верхнем регистре
	fmt.Println("В верхнем регистре:", strings.ToUpper(input))
}
