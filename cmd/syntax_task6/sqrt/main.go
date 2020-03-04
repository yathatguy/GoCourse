package main

import (
	"errors"
	"log"
	"math"

	"github.com/yathatguy/GoCourse/internal/numInput"
)

func main() {
	a := numInput.EnterIntNumber("Введите коэффициент перед X: ")
	b := numInput.EnterIntNumber("Введите коэффициент перед Y: ")
	c := numInput.EnterIntNumber("Введите константу: ")
	log.Printf("Получилось уравнение вида: %d*x^2 + %d*x + %d = 0", a, b, c)

	x1, x2, err := calc(a, b, c)
	if err != nil {
		log.Fatal(err)
	}

	if x1 != x2 {
		log.Printf("Выражение имеет два корня: %.2f и %.2f", x1, x2)
	} else {
		log.Printf("Выражение имеет один корень: %.2f", x1)
	}

}

func calc(a int, b int, c int) (float64, float64, error) {
	var x1, x2 float64
	d := b*b - 4*a*c
	if d < 0 {
		log.Println("do not support complex numbers")
		return x1, x2, errors.New("discriminant is below 0")
	}

	switch d>0 {
	case true:
		x1 = (float64(-b) + math.Sqrt(float64(d))) / (2*float64(a))
		x2 = (float64(-b) - math.Sqrt(float64(d))) / (2*float64(a))
 	case false:
		x1 = float64(-b) / (2*float64(a))
		x2 = x1
	}

	return x1, x2, nil
}
