package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	//values
	fmt.Println("go" + "lang" + ".org")
	fmt.Println("1+2", 1+2)
	fmt.Println(!(true || false && (true || false)))
	fmt.Println("float", 3.0/2.0)

	//variables
	var str string = "sample"
	fmt.Println(str)

	var a string
	fmt.Println(a)

	var b, c int = 2, 3
	fmt.Println(b, c, "sdfsd", "sdfds", "check", a)

	f := "short"
	e := 1
	d := false
	fmt.Println(d, e, f)
}
