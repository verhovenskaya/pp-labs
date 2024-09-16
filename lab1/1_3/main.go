//4)Написать программу для выполнения арифметических операций с двумя целыми числами и выводом результатов.

package main

import (
	"fmt"
	"os"
)

func main() {
	var k, j int

	fmt.Print("Введите первое значение: ")
	fmt.Fscan(os.Stdin, &k)

	fmt.Print("Введите второе значение: ")
	fmt.Fscan(os.Stdin, &j)

	sum := k + j
	diff := k - j
	mult := k * j
	div := k / j
	mod := k % j

	fmt.Println("Сумма:", sum)
	fmt.Println("Разность:", diff)
	fmt.Println("Произведение:", mult)
	fmt.Println("Частное:", div)
	fmt.Println("Остаток от деления:", mod)
}
