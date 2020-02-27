package main

import (
	"fmt"
	"log"
	"strconv"
)

// Написать функцию, которая будет получать позицию коня на шахматной доске,
// а возвращать массив из точек, в которые конь сможет сделать ход.
// Точку следует обозначить как структуру, содержащую x и y типа int
// Получение значений и их запись в точку должны происходить только с помощью отдельных методов.
// В них надо проводить проверку на то, что такая точка может существовать на шахматной доске.

var coordMap = map[int]string{
	1: "a",
	2: "b",
	3: "c",
	4: "d",
	5: "e",
	6: "f",
	7: "g",
	8: "h",
}

type Point struct {
	x int
	y int
}

func (p *Point) GetX() int {
	return p.x
}

func (p *Point) GetY() int {
	return p.y
}

func (p Point) String() string {
	return fmt.Sprintf("%s%d", coordMap[p.x], p.y)
}

func NewPoint(x, y int) Point {
	return Point{
		x: x,
		y: y,
	}
}

var availableMoves = []Point{
	NewPoint(1, 2),
	NewPoint(-1, 2),
	NewPoint(2, 1),
	NewPoint(-2, 1),
	NewPoint(1, -2),
	NewPoint(-1, -2),
	NewPoint(-2, -1),
	NewPoint(2, -1),
}

type Knight struct {
	position Point
}

func (k *Knight) GetPosition() *Point {
	return &k.position
}

func (k *Knight) CalculateAvailable() []Point {
	res := make([]Point, 0, 4)
	for _, v := range availableMoves {
		newX := k.GetPosition().GetX() + v.GetX()
		newY := k.GetPosition().GetY() + v.GetY()
		if isCoordAvailable(newX, newY) {
			res = append(res, NewPoint(newX, newY))
		}
	}

	return res
}

func isCoordAvailable(x, y int) bool {
	return x >= 1 && x <= 8 && y >= 1 && y <= 8
}

func NewKnight(p Point) Knight {
	return Knight{position: p}
}

func main() {
	var (
		x string
		y string
	)
	if _, err := fmt.Scanln(&x, &y); err != nil {
		fmt.Println(err)
	}

	startPoint := getStartPoint(x, y)


	kn := NewKnight(startPoint)
	fmt.Println(kn.CalculateAvailable())
}

func getStartPoint(x, y string) Point {

	startX, err := strconv.Atoi(x)
	if err != nil {
		log.Fatalf("Неверна задана стартовая точка, координата x: %v", err)
	}
	startY, err := strconv.Atoi(y)
	if err != nil {
		log.Fatalf("Неверна задана стартовая точка, координата y: %v", err)
	}

	startPoint := Point{
		x: startX,
		y: startY,
	}

	return startPoint
}
