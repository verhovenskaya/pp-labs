// Создать структуру Rectangle и реализовать метод для вычисления площади прямоугольника.
package main

import (
	"fmt"
	"os"
)

// Структура Rectangle, представляющая прямоугольник
type Rectangle struct {
	Width  int
	Height int
}

// Метод Area для вычисления площади прямоугольника
func (r Rectangle) Area() int {
	return r.Width * r.Height
}

func main() {
	var width, height int

	// Ввод данных от пользователя
	fmt.Println("Введите ширину прямоугольника:")
	fmt.Fscan(os.Stdin, &width)

	fmt.Println("Введите высоту прямоугольника:")
	fmt.Fscan(os.Stdin, &height)

	// Создание экземпляра структуры Rectangle
	rect := Rectangle{Width: width, Height: height}

	// Вычисление площади и вывод результата
	area := rect.Area()
	fmt.Println("Площадь прямоугольника:", area)
}
