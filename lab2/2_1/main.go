package main

import (
	"fmt"
	"os"
)

//Написать программу, которая определяет, является ли введенное пользователем число четным или нечетным.
//Реализовать функцию, которая принимает число и возвращает "Positive", "Negative" или "Zero".

func main() {
	var a int
	fmt.Println("Введите число а:")
	fmt.Fscan(os.Stdin, &a)
	fmt.Println(proverka(a))

}

func proverka(a int) string {
	if a%2 == 0 {
		return "chytnoe"
	} else {
		return "nechyot"
	}
}
