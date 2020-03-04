package statistics

func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

func Sum(xs []float64) float64 {
	sum := float64(0)
	for _, x := range xs {
		sum += x
	}
	return sum
}