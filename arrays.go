package main

import "fmt"

func main() {
	var x [5]float64
	x[4] = 100
	fmt.Println(len(x))
	total := 0.0
	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	fmt.Println(total/float64(len(x)))

	//shorter format
	a := [3]int{1, 2, 3}
	i := 0
	for i < len(a){
		fmt.Println(a[i])
		i++
	}

	//slices
	slice := make([]int, 5)
	fmt.Println(slice)

	slice1 := a[0:2]
	fmt.Println(slice1)

	//slice functions
	slice2 := append(slice1, 3, 4, 5)
	fmt.Println(slice2)

	//maps(hash)
	hash := make(map[string]int)
	hash["a"] = 2
	name, ok := hash["v"]
	fmt.Println(name, ok)
	hash["b"] = 2
	hash["c"] = 2
	delete(hash, "a")
	fmt.Println(hash)
}
