package main

import (
	"fmt"
	"os"
)

var (
	startAmount float64
	percentage float64
	finalSum float64
)

func main() {
	var length uint = 5
	getStartAmount()
	getPersentage()
	finalSum = calcFinalSum(startAmount, length)
	fmt.Printf("Через %.d лет выполучите %.2f рублей.", length, finalSum)
}

func getStartAmount() {
	fmt.Println("Сколько денег у вас есть?")
	if _, err := fmt.Scan(&startAmount); err != nil {
		fmt.Println("Введено некоректное float64 число")
		os.Exit(1)
	}
}

func getPersentage() {
	fmt.Println("Под какой процент кладем деньги в банк?")
	if _, err := fmt.Scan(&percentage); err != nil {
		fmt.Println("Введено некоректное float64 число")
		os.Exit(1)
	}
}

func calcFinalSum(amount float64, l uint) float64 {
	var total float64
	if l == 1 {
		total = startAmount
	} else {
		total = calcFinalSum(amount, l-1) * (1 + percentage / 100)
	}
	fmt.Println(total)
	return total
}