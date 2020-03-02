package main

import (
	"fmt"
	"sort"
)

type Contact struct {
	Name  string
	Phone []uint
}

type Book []Contact

func (b Book) Len() int {
	return len(b)
}

func (b Book) Less(i, j int) bool {
	return b[i].Name < b[j].Name
}

func (b Book) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func createAddressBook() Book {
	var addressBook []Contact

	alex := Contact{
		Name:  "Alex",
		Phone: []uint{89161234567},
	}
	addressBook = append(addressBook, alex)
	zug := Contact{
		Name:  "Zug",
		Phone: nil,
	}
	addressBook = append(addressBook, zug)
	bob := Contact{
		Name:  "Bob",
		Phone: []uint{89161234568, 89161234569},
	}
	addressBook = append(addressBook, bob)
	alice := Contact{
		Name:  "Alice",
		Phone: []uint{89161234570},
	}
	addressBook = append(addressBook, alice)

	return addressBook
}

func main() {
	book := createAddressBook()
	fmt.Println(book)
	sort.Sort(book)
	fmt.Println(book)
}
