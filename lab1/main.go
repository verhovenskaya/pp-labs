package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	//1) Написать программу, которая выводит текущее время и дату.
	currentTime := time.Now()
	fmt.Println("Time and Date:", currentTime.Format("2006-01-02 15:04:05"))

	//2)Создать переменные различных типов (int, float64, string, bool) и вывести их на экран.
	//3)Использовать краткую форму объявления переменных для создания и вывода переменных.

	var (
		a int     = 6
		b float64 = 6.6
		c string  = "Hello, World!"
		d bool    = true
	)

	// Вывод переменных на экран
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	//4)Написать программу для выполнения арифметических операций с двумя целыми числами и выводом результатов.
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

	//5)Реализовать функцию для вычисления суммы и разности двух чисел с плавающей запятой.

	var h, g float64

	fmt.Print("Первое значение с плавающей точкой:")
	fmt.Fscan(os.Stdin, &h)
	fmt.Print("Второе значение с плавающей точкой:")
	fmt.Fscan(os.Stdin, &g)
	sumFloat, diffFloat := calculateFloat(h, g)
	fmt.Println("Сумма чисел с плавающей запятой:", sumFloat)
	fmt.Println("Разность чисел с плавающей запятой:", diffFloat)

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

// Функция для вычисления суммы и разности двух чисел с плавающей запятой
func calculateFloat(a, b float64) (float64, float64) {
	return a + b, b - a
}

// Функция для вычисления среднего значения трех чисел
func average(a, b, c float64) float64 {
	return (a + b + c) / 3
}
