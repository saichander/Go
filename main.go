package main

import (
	"GO_WorkSpace/math"
	"fmt"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	fmt.Println("Average of given slice", math.Average(xs))
}
