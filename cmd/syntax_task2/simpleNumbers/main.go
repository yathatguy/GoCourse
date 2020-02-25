package main

import (
	"fmt"
	"os"
)

func main() {
	arr := createArray(1000)
	simpleArr := getSimpleArray(arr)
	fmt.Println(simpleArr)
}

func createArray(l int) []int {
	resArr := make([]int, l, l)
	if l < 1 {
		fmt.Printf("  unable to create array with %v in length\n", l)
		os.Exit(1)
	}
	for i := 1; i <= l; i++ {
		resArr[i] = i
	}
	return resArr
}

func getSimpleArray(arr []int) []int {
	simpleArr := make([]int, len(arr), len(arr))
	for i := 1; i < len(arr); i++ {
		for j := i+1; j < len(arr); j++ {
			if arr[j] != 0 && arr[i] != 0{
				if arr[j] % arr[i] == 0 {
					arr[j] = 0
				}
			}
		}
	}
	var j = 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			simpleArr[j] = arr[i]
			j += 1
		}
	}
	// TODO: resArr should not be fixed length, it can be constructed with blocks, each new block adds if getSimpleArray is not fit to target length
	return simpleArr[:100]
}
