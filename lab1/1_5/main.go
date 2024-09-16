package main

import (
	"fmt"
	"os"
)

// Функция для вычисления среднего значения трех чисел
func average(a, b, c float64) float64 {
	return (a + b + c) / 3
}

func main() {
	//6)Написать программу, которая вычисляет среднее значение трех чисел.
	var m, n, l float64
	fmt.Println("Введите первое число:")
	fmt.Fscan(os.Stdin, &m)
	fmt.Println("Введите второе число:")
	fmt.Fscan(os.Stdin, &n)
	fmt.Println("Введите третье число:")
	fmt.Fscan(os.Stdin, &l)
	fmt.Println("Среднее значение:", average(m, n, l))

}
