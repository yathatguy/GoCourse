package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	addressBook := make(map[string][]int)

	addressBook["Alex"] = []int{89996543210}
	addressBook["Bob"] = []int{89167243812}
	addressBook["Bob"] = append(addressBook["Bob"], 89155243627)

	fmt.Println(addressBook)

	for name, numbers := range addressBook {
		fmt.Println("Абонент:", name)
		for i, number := range numbers {
			fmt.Printf("\t %v) %v \n", i+1, number)
		}
	}

	addBookJson, err := json.Marshal(addressBook)
	if err != nil {
		fmt.Printf("Ошибка конвертации в json: %v\n", err)
	}

	err = ioutil.WriteFile("addBookJson.json", addBookJson, 0644)
	if err != nil {
		fmt.Printf("Ошибка записи в файл: %v\n", err)
	}
}
