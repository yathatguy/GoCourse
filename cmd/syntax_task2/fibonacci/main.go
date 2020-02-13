package main

import (
	"fmt"
	"math/big"
)

func main() {
	fibRow := make([]*big.Int, 100, 100)
	for i := 0; i<100; i++ {
		switch i {
		case 0: fibRow[i] = big.NewInt(0)
		case 1: fibRow[i] = big.NewInt(1)
		default: fibRow[i] = big.NewInt(0).Add(fibRow[i-1], fibRow[i-2])
		}
		fmt.Printf("%v число в последовательности Фибоначчи: %v\n", i+1, fibRow[i])
	}
}
