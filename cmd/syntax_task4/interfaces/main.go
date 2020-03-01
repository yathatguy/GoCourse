package main

import "fmt"

type Human struct {
	Name   string
	Age    uint
	Gender string
	Alive  bool
	Weight float32
	Height float32
}

func (h Human) age() uint {
	return h.Age
}

func (h Human) height() float32 {
	return h.Height
}

func (h Human) name() string {
	return h.Name
}

type Animal struct {
	Name   string
	Age    uint
	Type   string
	Alive  bool
	Weight float32
	Height float32
}

func (a Animal) age() uint {
	return a.Age
}

func (a Animal) height() float32 {
	return a.Height
}

func (a Animal) name() string {
	return a.Name
}

type Furniture struct {
	Brand  string
	Type   string
	Weight float32
	Height float32
	Depth  float32
}

func (f Furniture) height() float32 {
	return f.Height
}

func (f Furniture) name() string {
	return f.Brand
}

type Creature interface {
	name() string
	age() uint
	height() float32
}

func age(creature Creature) uint {
	return creature.age()
}

type Thing interface {
	name() string
	height() float32
}

func height(thing Thing) float32 {
	return thing.height()
}

func name(thing Thing) string {
	return thing.name()
}

func main() {
	man := Human{
		Name:   "Dima",
		Age:    36,
		Gender: "M",
		Alive:  true,
		Weight: 60.5,
		Height: 168.0,
	}
	dog := Animal{
		Name:   "Fry",
		Age:    7,
		Type:   "Dog",
		Alive:  true,
		Weight: 10.0,
		Height: 37,
	}
	sofa := Furniture{
		Brand:  "Admiral",
		Type:   "Sofa",
		Weight: 32.5,
		Height: 40,
		Depth:  120,
	}

	for _, c := range []Creature{man, dog} {
		fmt.Printf("The %s's age is: %d\n", name(c), age(c))
	}

	for _, t := range []Thing{man, dog, sofa} {
		fmt.Printf("The %s's height is: %.2f\n", name(t), height(t))
	}
}
