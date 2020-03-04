package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

const (
	distance = 500
	maxSpeed = 180
	numCars  = 3
)

type Car struct {
	name  string
	speed uint
}

var (
	cars []Car
	wg   sync.WaitGroup
	name string
)
var finish = make(chan string)

func main() {
	// Prepare cars
	for i := 0; i < numCars; i++ {
		wg.Add(1)
		go createCar(strconv.Itoa(i))
	}
	wg.Wait()

	// Start race

	for _, car := range cars {
		go race(car)
	}
	res := <-finish
	time.Sleep(1 * time.Second)
	fmt.Printf("Car %s has won!", res)
}

func createCar(s string) {
	speed := uint(rand.Float32() * maxSpeed)
	cars = append(cars, Car{
		name:  s,
		speed: speed,
	})
	wg.Done()
	return
}

func race(car Car) {
	var point uint
	for {
		point += car.speed
		time.Sleep(1 * time.Second)
		fmt.Printf("Car %s, distance %d\n", car.name, point)
		if point >= distance {
			finish <- car.name
			return
		}
	}
}
