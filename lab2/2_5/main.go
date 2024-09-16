// Создать структуру Rectangle и реализовать метод для вычисления площади прямоугольника.
package main

import (
	"fmt"
	"os"
)

func main() {
	var a int
	var b int
	var c int
	fmt.Println("Введите а:")
	fmt.Fscan(os.Stdin, &a)
	fmt.Println("Введите b:")
	fmt.Fscan(os.Stdin, &b)
	fmt.Println("Введите c:")
	fmt.Fscan(os.Stdin, &c)
	fmt.Println(square(a, b, c))

}
func square(a, b, c int) int {
	s := a * b * c
	return s
}
