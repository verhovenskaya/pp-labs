//Написать функцию, которая принимает два целых числа и возвращает их среднее значение.

package main

import (
	"fmt"
	"os"
)

func main() {
	var a, b int
	fmt.Println("")
	fmt.Fscan(os.Stdin, &a)
	fmt.Println("")
	fmt.Fscan(os.Stdin, &b)
	fmt.Println(average(a, b))

}

func average(a, b int) float64 {
	return float64(a+b) / 2
}
