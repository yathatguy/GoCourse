package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	spinner(50 * time.Millisecond)
	wg.Wait()
}

func spinner(delay time.Duration) {
	done := time.After(10 * time.Second)
	defer wg.Done()
	for {
		for _, r := range "-\\|/" {
			fmt.Printf("%c\r", r)
			time.Sleep(delay)
		}

		select {
		case <-done:
			return
		default:
			continue
		}
	}
}
