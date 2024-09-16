// Написать функцию, которая принимает строку и возвращает ее длину.
package main

import (
	"fmt"
	"os"
)

func main() {
	var a string
	fmt.Fscan(os.Stdin, &a)
	fmt.Println(len(a))
}
func length(a string) int {
	return len(a)
}
