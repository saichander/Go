package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

func distance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1))
}

//methods
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

//embedded types
type Person struct {
	Name string
	age  int
}

type Android struct {
	Person
	model string
}

func (p *Person) talk() {
	fmt.Println("Hi my name is", p.Name, "and age is", p.age)
}

//interfaces
type Shape interface {
	area() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func main() {
	//structs
	fmt.Println("Structs")

	c := Circle{x: 2, y: 3, r: 5}
	fmt.Println(c.x, c.y, c.r) // 2 3 5

	fmt.Println("Area of circle using structs", circleArea(c))
	fmt.Println("Area of circle using methods", c.area()) //go automatically sends c as pointer to circle for this method

	r := Rectangle{x1: 1, y1: 2, x2: 3, y2: 4}
	fmt.Println("Area of rectangle", r.area())

	p := Person{"test", 10}
	p.talk()

	a := Android{p, "test_model"}
	fmt.Println(a.Person.Name)
	a.talk() //we can directly call person methods on android

	fmt.Println(totalArea(&c, &r))
}
