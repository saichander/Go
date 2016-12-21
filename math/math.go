package math

func Average(xs []float64) (total float64) {
	for _, s := range xs {
		total += s
	}
	total += total / float64(len(xs))
}
