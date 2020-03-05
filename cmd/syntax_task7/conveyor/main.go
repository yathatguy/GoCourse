package main

import (
	"fmt"
)

func main() {
	n := 5
	naturals := make(chan int)
	squares := make(chan int)

	// генерация
	go func() {
		for x := 0; ; x++ {
			naturals <- x
			if x >= n {
				close(naturals)
				break
			}
		}
	}()

	// возведение в квадрат
	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				close(squares)
				return
			}
			squares <- x * x
		}
	}()

	// печать
	for {
		res, ok := <-squares
		if !ok {
			break
		}
		fmt.Println(res)
	}
}
