package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// Set up flags
	pattern := flag.String("pattern", "*", "pattern to search for in file")
	file := flag.String("file", "", "file name where search for pattern will be performed")
	insense := flag.Bool("i", false, "perform case insensitive search")

	// Parse flags
	flag.Parse()

	// Check input params
	if *file == "" {
		log.Fatal("file name was not provided")
	}

	// Open source file for reading
	src, err := os.OpenFile(*file, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatalf("can't open file: %v", err)
	}
	defer src.Close()

	f := bufio.NewReader(src)
	for {
		l, pref, err := f.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("cannot read new line from file: %v", err)
		}
		searchL := l
		if pref {
			fmt.Println("line it too long, no search was performed")
			continue
		}
		if *insense {
			searchL = []byte(strings.ToLower(string(l)))
		}
		if err != nil {
			log.Printf("cannot read new line from file: %v", err)
		}
		if *pattern == "*" {
			fmt.Println(string(l))
		}
		if strings.Contains(string(searchL), *pattern) {
			fmt.Println(string(l))
		}
	}
}
