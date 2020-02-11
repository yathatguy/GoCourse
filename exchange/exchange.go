package main

import (
	"fmt"
	"os"
)

func main() {
	const rubToUsdCourse float64 = 62.34
	var rubAmount float64
	fmt.Println("Сколько рублей у вас есть?")
	if _, err := fmt.Scanln(&rubAmount); err != nil {
		fmt.Println("Вы ввели не число...")
		os.Exit(1)
	}
	fmt.Printf("У вас есть %.2f рублей\n", rubAmount)
	fmt.Printf("В долларах это стоставляет: %.2f долларов\n", rubAmount/rubToUsdCourse)
}
