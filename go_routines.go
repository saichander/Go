package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(n int) {
	fmt.Println("inside go routine func")
	for i := 0; i < 3; i++ {
		fmt.Println(n, ":", i)
	}
}

func main() {
	amt := time.Duration(rand.Intn(250))
	for i := 0; i < 3; i++ {
		go f(i)
		time.Sleep(time.Millisecond * amt)
	}
	fmt.Println("after calling go routine func")
	var input string
	fmt.Scanln(&input)
}
