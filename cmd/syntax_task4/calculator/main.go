package main

import (
	"fmt"

	"github.com/yathatguy/GoCourse/internal/calculator"
)

func main() {
	input := ""
	for {
		fmt.Print("> ")
		if _, err := fmt.Scanln(&input); err != nil {
			fmt.Println(err)
			continue
		}

		if input == "exit" {
			break
		}

		if input == "help" {
			printHelp()
			continue
		}

		if res, err := calculator.Calculate(input); err == nil {
			fmt.Printf("Результат: %v\n", res)
		} else {
			fmt.Println("Не удалось произвести вычисление")
		}
	}
}

func printHelp() {
	fmt.Println("Тут должно быть описание help для программы \"Калькулятор\"")
}
