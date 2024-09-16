//Написать программу, которая определяет, является ли введенное пользователем число четным или нечетным.
//Реализовать функцию, которая принимает число и возвращает "Positive", "Negative" или "Zero".
//Написать программу, которая выводит все числа от 1 до 10 с помощью цикла for.
//Написать функцию, которая принимает строку и возвращает ее длину.
//Создать структуру Rectangle и реализовать метод для вычисления площади прямоугольника.
//Написать функцию, которая принимает два целых числа и возвращает их среднее значение.

package main

import (
	"fmt"
)

// Функция, которая определяет, является ли число четным или нечетным
func isEvenOrOdd(num int) string {
	if num%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

// Функция, которая возвращает "Positive", "Negative" или "Zero"
func checkNumber(num int) string {
	if num > 0 {
		return "Positive"
	} else if num < 0 {
		return "Negative"
	} else {
		return "Zero"
	}
}

// Функция, которая возвращает длину строки
func stringLength(s string) int {
	return len(s)
}

// Структура Rectangle
type Rectangle struct {
	Width, Height float64
}

// Метод для вычисления площади прямоугольника
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Функция, которая возвращает среднее значение двух целых чисел
func average(a, b int) float64 {
	return float64(a+b) / 2
}

func main() {
	// Проверка четности и нечетности
	var num int
	fmt.Print("Введите число: ")
	fmt.Scan(&num)
	fmt.Printf("Число %d является %s\n", num, isEvenOrOdd(num))

	// Проверка положительности, отрицательности или нуля
	fmt.Printf("Число %d является %s\n", num, checkNumber(num))

	// Вывод чисел от 1 до 10
	fmt.Println("Числа от 1 до 10:")
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// Проверка длины строки
	var str string
	fmt.Print("Введите строку: ")
	fmt.Scan(&str)
	fmt.Printf("Длина строки \"%s\" составляет %d\n", str, stringLength(str))

	// Вычисление площади прямоугольника
	var width, height float64
	fmt.Print("Введите ширину прямоугольника: ")
	fmt.Scan(&width)
	fmt.Print("Введите высоту прямоугольника: ")
	fmt.Scan(&height)
	rect := Rectangle{Width: width, Height: height}
	fmt.Printf("Площадь прямоугольника составляет %.2f\n", rect.Area())

	// Вычисление среднего значения
	var a, b int
	fmt.Print("Введите два целых числа: ")
	fmt.Scan(&a, &b)
	fmt.Printf("Среднее значение чисел %d и %d составляет %.2f\n", a, b, average(a, b))
}
