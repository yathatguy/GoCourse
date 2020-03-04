package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	cancel := make(chan int)
	go func() {
		bs := make([]byte, 4)
		if _, err := os.Stdin.Read(bs); err != nil {
			log.Println(err)
		}
		if string(bs) == "exit" {
			cancel <- 1
		}
	}()

	fmt.Println("Countdown started. Print \"exit\" and hit [Enter] to cancel")
	tick := make(<-chan time.Time)
	tick = time.Tick(1 * time.Second)

	for countdown := 10; countdown > 0; countdown-- {
		select {
		case moment := <-tick:
			fmt.Println(moment.Format("15:04:05"), countdown)
		case <-cancel:
			fmt.Println("Launch canceled!")
			return
		}
	}
	fmt.Println("The rocket starts!")
}