package main

import (
	"GoCourse/pkg/mod/golang.org/x/tools@v0.0.0-20191030062658-86caa796c7ab/go/ssa/interp/testdata/src/os"
	"fmt"
	"math"
)

func main() {

	c1 := getCathetus("Введите длину первого катета в float64")
	c2 := getCathetus("Введите длину второго катета в float64")

	square := calcSquare(c1, c2)
	fmt.Println("Площадь треугольника: ", square)
	parimeter := calcPerimeter(c1, c2)
	fmt.Println("Периметр равен: ", parimeter)
	hypotenuse := calcHypotenuse(c1, c2)
	fmt.Println("Гипотенуза равна: ", hypotenuse)
}

func getCathetus(s string) float64 {
	var c float64
	fmt.Println(s)
	if _, err := fmt.Scan(&c); err != nil {
		fmt.Println("Введено некоректное значение")
		os.Exit(1)
	}
	return c
}

func calcSquare(c1 float64, c2 float64) float64 {
	return 0.5 * c1 * c2
}

func calcPerimeter(c1 float64, c2 float64) float64 {
	return c1 + c2 + calcHypotenuse(c1, c2)
}

func calcHypotenuse(c1 float64, c2 float64) float64 {
	return math.Sqrt(math.Pow(c1, 2) + math.Pow(c2, 2))
}
