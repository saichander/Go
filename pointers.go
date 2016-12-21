package main

import "fmt"

func zero(xPtr *int) {
	fmt.Println(*xPtr)
	*xPtr = 0
}

func swap(x *int, y *int) {
	*x = *x + *y
	*y = *x - *y
	*x = *x - *y
}
func main() {
	i := 2
	zero(&i)
	fmt.Println(i)

	//new
	a := new(int)
	fmt.Println("created by new", a)
	zero(a)
	fmt.Println("value of a", *a)
	x := 2
	y := 3
	swap(&x, &y)
	fmt.Println(x, y)
}
