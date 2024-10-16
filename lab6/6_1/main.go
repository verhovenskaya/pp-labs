package main

import (
	"fmt"
	"math/rand"
	"time"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func generateRandomNumbers(count int) []int {
	numbers := make([]int, count)
	for i := 0; i < count; i++ {
		numbers[i] = rand.Intn(100)
	}
	return numbers
}

func sumSeries(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func main() {
	for i := range 10 {
		go factorial(i)
		go generateRandomNumbers(i)
		go sumSeries(i)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Все горутины завершены")
}
