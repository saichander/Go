package math

import "testing"

type testpair struct {
	values  []float64
	average float64
}

var tests = []testpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 2, 3, 4}, 2.5},
	{[]float64{333, 333, 333}, 333},
}

func TestAverage(t *testing.T) {
	for _, pair := range tests {
		var v float64
		v = Average(pair.values)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"Got", v)
		}
	}
}
