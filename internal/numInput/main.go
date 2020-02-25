package numInput

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func inputData(msg string) (data string){
	fmt.Print(msg)
	if _, err := fmt.Scanln(&data); err != nil {
		log.Println("input error: ", err)
	}
	if data == "exit" {
		os.Exit(0)
	}
	return
}

func EnterIntNumber(msg string) int {
	data := inputData(msg)
	num, err := strconv.ParseInt(data, 10, 32)
	if err != nil {
		log.Fatalf("unable to parse input: %s", err)
	}
	return int(num)
}

func EnterFloat32Number(msg string) float32 {
	data := inputData(msg)
	num, err := strconv.ParseFloat(data, 32)
	if err != nil {
		log.Fatalf("unable to parse input: %s", err)
	}
	return float32(num)
}

func EnterFloat64Number(msg string) float64 {
	data := inputData(msg)
	num, err := strconv.ParseFloat(data, 64)
	if err != nil {
		log.Fatalf("unable to parse input: %s", err)
	}
	return num
}