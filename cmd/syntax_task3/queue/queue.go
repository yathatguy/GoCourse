package queue

import "errors"

// FIFO
var queue []int

func Push(i int) {
	queue = append(queue, i)

}

func Pop() (int, error) {
	if len(queue) == 0 {
		return 0, errors.New("queue is empty")
	}

	val := queue[0]
	for i:= 0; i < len(queue)-1; i++ {
		queue[i] = queue[i+1]
	}
	queue = queue[:len(queue)-1]
	return val, nil
}
