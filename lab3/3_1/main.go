package main

import (
	"fmt"
	"os"
)

func main() {
	var a int
	fmt.Println()
	fmt.Fscan(os.Stdin, &a)
	fmt.Println(mathutils.fact(a))
}
