package main

import "fmt"


func avg(a []float64) float64{
	total := 0.0
	for _, v := range a {
		total += v
	}
	return total / float64(len(a))
}

func test() string{
return "test"
}

//variadic functions
func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func makeEvengenerator() func() uint{
	i :=  uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x*factorial(x-1)
}



func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}

func third() {
	fmt.Println("3rd")
}

func main() {
fmt.Println("Functions")
a := []float64{ 12, 23, 23, 21, 12}
fmt.Println("average of array", avg(a))
fmt.Println(test())
fmt.Println("sum of 1,2,3", add(1,2,3))

//closure
nextEven := makeEvengenerator() //returns a function which retuns next even number
fmt.Println(nextEven()) // 0
fmt.Println(nextEven()) // 2
fmt.Println(nextEven()) // 4

fmt.Println("factorial of 10", factorial(10)) //prints factorial of the given number

//using defer
defer first() //defer functions are stored in stack, last defer will be executed first
defer second()
third()
defer func() {
	str := recover()
	fmt.Println(str)
}()
panic("Don't panic we have recover function")
}
