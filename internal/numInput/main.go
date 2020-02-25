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
	for {
		num, err := strconv.ParseFloat(inputData(msg), 32)
		if err == nil {
			// Ввод числа прошел без ошибок
			return int(num)
		}
		log.Fatalln("unable to parse input")
	}
}

func EnterFloat32Number(msg string) float32 {
	for {
		num, err := strconv.ParseFloat(inputData(msg), 32)
		if err == nil {
			// Ввод числа прошел без ошибок
			return float32(num)
		}
		log.Fatalln("unable to parse input")
	}
}

func EnterFloat64Number(msg string) float64 {
	for {
		num, err := strconv.ParseFloat(inputData(msg), 32)
		if err == nil {
			// Ввод числа прошел без ошибок
			return num
		}
		log.Fatalln("unable to parse input")
	}
}