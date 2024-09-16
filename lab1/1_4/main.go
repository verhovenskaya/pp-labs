// 5)Реализовать функцию для вычисления суммы и разности двух чисел с плавающей запятой.
package main

import (
	"fmt"
	"os"
)

func main() {
	var h, g float64

	fmt.Print("Первое значение с плавающей точкой:")
	fmt.Fscan(os.Stdin, &h)
	fmt.Print("Второе значение с плавающей точкой:")
	fmt.Fscan(os.Stdin, &g)
	sumFloat, diffFloat := calculateFloat(h, g)
	fmt.Println("Сумма чисел с плавающей запятой:", sumFloat)
	fmt.Println("Разность чисел с плавающей запятой:", diffFloat)
}

// Функция для вычисления суммы и разности двух чисел с плавающей запятой
func calculateFloat(a, b float64) (float64, float64) {
	return a + b, b - a
}
