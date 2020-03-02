package main

import (
	"fmt"

	"github.com/yathatguy/GoCourse/internal/numInput"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	num := numInput.EnterIntNumber("Введите число: ")
	if num%2 == 0 {
		fmt.Println("Четное")
	} else {
		fmt.Println("Не четное")
	}
}
