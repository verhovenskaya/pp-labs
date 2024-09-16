package main

import (
	"fmt"
	"time"
)

func main() {
	//1) Написать программу, которая выводит текущее время и дату.
	currentTime := time.Now()
	fmt.Println("Time and Date:", currentTime.Format("2006-01-02 15:04:05"))
}
