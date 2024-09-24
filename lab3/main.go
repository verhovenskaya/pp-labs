package main

import (
	"fmt"
	"pp-labs/lab3/mathutils"
	"pp-labs/lab3/stringutils"
)

func main() {
	//Написать программу, которая создает массив из 5 целых чисел, заполняет его значениями и выводит их на экран.
	fmt.Println("Введите число для вычисления факториала:")
	var num int
	fmt.Scanln(&num)
	factorial := mathutils.Factorial(num)
	fmt.Printf("Факториал %d равен %d\n", num, factorial)

	//Переворот строки
	fmt.Println("Введите строку для переворота:")
	var str string
	fmt.Scanln(&str)
	reversedStr := stringutils.Reverse(str)
	fmt.Printf("Перевернутая строка: %s\n", reversedStr)

	//Массив целых чисел
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Массив целых чисел:", numbers)

	//Срез из массива
	slice := numbers[1:4] //срез из 2-го по 4-ый элемент
	fmt.Println("Срез:", slice)

	//Добавление элементов в срез
	slice = append(slice, 6)
	fmt.Println("Срез после добавления:", slice)

	//Удаление элемента из среза ---
	slice = append(slice[:2], slice[3:]...) // удаляем 3-ий элемент
	fmt.Println("Срез после удаления:", slice)

	//Нахождение самой длинной строки
	strings := []string{"Hello", "World", "This", "is", "a", "long", "string"}
	longestString := findLongestString(strings)
	fmt.Printf("Самая длинная строка: %s\n", longestString)
}

// Функция для нахождения самой длинной строки в срезе
func findLongestString(strings []string) string {
	longest := strings[0]
	for _, str := range strings {
		if len(str) > len(longest) {
			longest = str
		}
	}
	return longest
}
