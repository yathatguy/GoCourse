package queue

import "errors"

// FIFO
var queue []int

func Push(i int) {
	queue = append(queue, i)

}

func Pop() ([]int, int, error) {
	if len(queue) == 0 {
		return nil, 0, errors.New("queue is empty")
	} else {
		val := queue[0]
		for i:= 0; i < len(queue)-1; i++ {
			queue[i] = queue[i+1]
		}
		return queue[:len(queue)-1], val, nil
	}
}
