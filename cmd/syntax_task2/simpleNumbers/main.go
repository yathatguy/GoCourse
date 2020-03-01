package main

import (
	"fmt"
	"math"
)

func main() {
	findSimple(100)
}

func findSimple(n int) {
	foundCnt := 0
	isDividable := false
	for i := 1; foundCnt < n; i++ {
		switch i {
		case 1:
			fmt.Println(1)
		default:
			isDividable = false
			for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
				if i%j == 0 {
					isDividable = true
					break
				}
			}
			if !isDividable {
				foundCnt++
				fmt.Println(i)
			}
		}
	}
}
