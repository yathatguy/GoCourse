package main

import (
	"fmt"
	"math/big"
)

func main() {
	var (
		prevMinusOne = big.NewInt(0)
		prev = big.NewInt(0)
		fibNum = big.NewInt(0)
	)
	for i := 0; i<100; i++ {
		switch i {
		case 0:
			fibNum = big.NewInt(0)
		case 1:
			fibNum = big.NewInt(1)
			prev = fibNum
		default:
			fibNum = big.NewInt(0).Add(prev, prevMinusOne)
			prevMinusOne = prev
			prev = fibNum
		}
		fmt.Printf("%v число в последовательности Фибоначчи: %v\n", i+1, fibNum)
	}
}
