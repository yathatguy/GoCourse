package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./cmd/syntax_task5/readLines/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	bsSize := 512
	bs := make([]byte, bsSize)
	blocks := int(stat.Size() / int64(bsSize))

	for i := 0; i <= blocks+1; i++ {
		a, err := file.Read(bs)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(a)
	}
}
