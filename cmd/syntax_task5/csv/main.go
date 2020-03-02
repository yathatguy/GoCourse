package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	recs := createRecords()

	// Write file
	f1, err := os.OpenFile("cmd/syntax_task5/csv/cvsText.cvs", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("can't open file: %v", err)
	}
	defer f1.Close()

	w := csv.NewWriter(f1)
	if err := w.WriteAll(recs); err != nil {
		log.Fatalln("error writing csv file:", err)
	}
	fmt.Println(recs)

	// Read file
	f2, err := os.OpenFile("cmd/syntax_task5/csv/cvsText.cvs", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatalf("can't open file: %v", err)
	}
	defer f2.Close()
	r := csv.NewReader(f2)
	res, err := r.ReadAll()
	if err != nil {
		log.Fatalln("error reading csv file:", err)
	}
	fmt.Println(res)

}

func createRecords() [][]string {
	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}
	return records
}
