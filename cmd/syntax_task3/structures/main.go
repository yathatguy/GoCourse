package main

import "fmt"

type Car struct {
	Brand string
	Model string
	Year string
	Trunk struct {
		Volume uint
		Fullness uint
	}
	Engine struct {
		Size float32
		Started bool
	}
}

func main() {
	merGle := createGle()
	Maz := createMaz()
	checkNewer(merGle, Maz)
}

func createGle() Car {
	car := Car{
		Brand: "Mercedes",
		Model: "GLE Coupe",
		Year:  "2018",
		Trunk: struct {
			Volume   uint
			Fullness uint
		}{
			Volume: 3000,
			Fullness: 2700,
		},
		Engine: struct {
			Size    float32
			Started bool
		}{
			Size: 3000,
			Started: false,
		},
	}
	return car
}

func createMaz() Car {
	car := Car{
		Brand: "МАЗ",
		Model: "5440",
		Year:  "2011",
		Trunk: struct {
			Volume   uint
			Fullness uint
		}{
			Volume: 0,
			Fullness: 0,
		},
		Engine: struct {
			Size    float32
			Started bool
		}{
			Size: 11100,
			Started: false,
		},
	}
	return car
}

func checkNewer(car1 Car, car2 Car) {
	if car1.Year < car2.Year {
		fmt.Printf("Авто %v %v старше %v %v", car1.Brand, car1.Model, car2.Brand, car2.Model)
	} else if car1.Year > car2.Year {
		fmt.Printf("Авто %v %v младше %v %v", car1.Brand, car1.Model, car2.Brand, car2.Model)
	} else {
		fmt.Println("Оба авто одинакового возраста")
	}
}
