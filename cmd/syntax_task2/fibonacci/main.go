package main

import "fmt"

func main() {
	fibRow := make([]float64, 100, 100)
	for i := 0; i<100; i++ {
		switch i {
		case 0: fibRow[i] = 0
		case 1: fibRow[i] = 1
		default: fibRow[i] = fibRow[i-1] + fibRow[i-2]
		}
		fmt.Printf("%v число в последовательности Фибоначчи: %v\n", i+1, fibRow[i])
	}
}
