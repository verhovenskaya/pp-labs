package main

import (
	"fmt"
	"os"
)

//Реализовать функцию, которая принимает число и возвращает "Positive", "Negative" или "Zero".

func main() {
	var a int
	fmt.Println(" Введите число а:")
	fmt.Fscan(os.Stdin, &a)
	fmt.Println(kakoyeChislo(a))
}
func kakoyeChislo(a int) string {
	if a > 0 {
		return "Positive"
	} else if a == 0 {
		return "Zero"
	} else {
		return "Negative"
	}
}
