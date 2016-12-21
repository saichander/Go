package main

import (
	"fmt"
)

func main() {
	c1 := make(chan int, 10)

	go func() {
		for i := 0; ; i++ {
			c1 <- i
		}
	}()

	go func() {
		for {
			select {
			case <-c1:
				fmt.Println(<-c1)
			default:
				fmt.Println("Nothing ready")
			}
		}
	}()

	var input int64
	fmt.Scanln(&input)
}
