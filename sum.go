package main

import "fmt"

func main() {
	// Задача: Найти сумму всех чисел от 1 до 100

	sum := 0

	for i := 1; i <= 100; i++ {
		sum += i
	}

	fmt.Println("Сумма чисел от 1 до 100:", sum)
}
