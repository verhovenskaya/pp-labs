package main

import (
	"fmt"
)

func main() {
	//Создать переменные различных типов (int, float64, string, bool) и вывести их на экран.
	//Использовать краткую форму объявления переменных для создания и вывода переменных.

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
}
