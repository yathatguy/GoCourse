package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	// Parse flags
	source := flag.String("src", "", "source file")
	dest := flag.String("dst", "", "destination")
	flag.Parse()

	// Check if all params are provided
	if *source == "" || *dest == "" {
		log.Fatal("not all required params are provided")
	}

	// Open source file
	s, err := os.OpenFile("./" + *source, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatalf("can't open file: %v", err)
	}
	defer s.Close()

	// Open destination file
	d, err := os.OpenFile("./" + *dest, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("can't open file: %v", err)
	}
	defer d.Close()

	// Copy data blocks from source to destination
	stat, err := s.Stat()
	if err != nil {
		log.Fatal(err)
	}

	bsSize := 512
	bs := make([]byte, bsSize)
	blocks := int(stat.Size() / int64(bsSize))

	for i := 0; i < blocks+1; i++ {
		l, err := s.Read(bs)
		if err != nil {
			log.Fatal(err)
		}
		_, err = d.Write(bs[:l])
		if err != nil {
			log.Fatal(err)
		}
	}
}
