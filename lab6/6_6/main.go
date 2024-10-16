package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Job struct {
	Content string
}

func ReverseString(s string) string {
	reversed := make([]byte, len(s))
	for i, j := 0, len(s)-1; i < len(s); i, j = i+1, j-1 {
		reversed[i] = s[j]
	}
	return string(reversed)
}

func processTasks(taskChannel <-chan Job) {
	for job := range taskChannel {
		result := ReverseString(job.Content)
		fmt.Println(result)
	}
}

func main() {
	var workerCount int
	fmt.Print("Укажите количество рабочих потоков: ")
	fmt.Scan(&workerCount)

	taskChannel := make(chan Job)

	for i := 0; i < workerCount; i++ {
		go processTasks(taskChannel)
	}

	file, err := os.Open("ку.txt")
	if err != nil {
		fmt.Println("Не удалось открыть файл:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		taskChannel <- Job{Content: scanner.Text()}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
	}

	close(taskChannel)
	time.Sleep(1 * time.Second)
}
