package main

import (
	"errors"
	"reflect"
	"testing"
)

type testpair struct {
	input  []int
	output []float64
	err    error
}

var tests = []testpair{
	{
		[]int{1, 0, -4},
		[]float64{2, -2},
		nil,
	},
	{
		[]int{1, 3, -4},
		[]float64{1, -4},
		nil,
	},
	{
		[]int{1, -5, 9},
		[]float64{2, 2},
		errors.New("discriminant is below 0"),
	},
	{
		[]int{1, -4, 4},
		[]float64{2, 2},
		nil,
	},
}

func TestCalc(t *testing.T) {
	for _, pair := range tests {
		x1, x2, err := calc(pair.input[0], pair.input[1], pair.input[2])
		if err != nil {
			if !reflect.DeepEqual(err, pair.err) {
				t.Errorf("%v does not equal %v", err, pair.err)
			}
			continue
		}
		if !reflect.DeepEqual([]float64{x1, x2}, pair.output) {
			t.Errorf(
				"For %d*x^2 + %d*x + %d expected %.2f and %.2f got %.2f and %.2f", pair.input[0], pair.input[1], pair.input[2],
				pair.output[0], pair.output[1],
				x1, x2,
			)
		}
	}
}
