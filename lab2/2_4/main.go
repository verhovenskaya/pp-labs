// Написать функцию, которая принимает строку и возвращает ее длину.
package main

import (
	"fmt"
)

func main() {
	var a string
	fmt.Scanln(&a)
	fmt.Print(length(a))
}
func length(a string) int {
	return len(a)
}
