package statistics

import "testing"

func TestAverage(t *testing.T) {
	var v float64
	v = Average([]float64{1, 2})
	if v != 1.5 {
		t.Error("Expected 1.5, got ", v)
	}
}

type pairs struct {
	input []float64
	res float64
}
var testData = []pairs{
	{
		[]float64{1, 2, 3},
		float64(6),
	},
	{
		[]float64{1, 0},
		float64(1),
	},
	{
		[]float64{1.0, -7, 3, 22.2},
		float64(19.2),
	},
}

func TestSum(t *testing.T) {
	for _, pair := range testData {
		v := Sum(pair.input)
		if v != pair.res {
			t.Error(
				"For", pair.input,
				"expected", pair.res,
				"got", v,
			)
		}
	}
}
