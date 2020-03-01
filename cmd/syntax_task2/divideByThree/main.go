package main

import (
	"fmt"

	"github.com/yathatguy/GoCourse/internal/numInput"
)

func main() {
	num := numInput.EnterIntNumber("Введите число: ")
	if num%3 == 0 {
		fmt.Printf("Число %v делится без остатка на 3", num)
	} else {
		fmt.Printf("Число %v не делится без остатка на 3", num)
	}
}
