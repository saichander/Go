package main

import (
	"fmt"
	"time"
)

func pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func ponger(c chan<- string) { // <- on right side of chan indicates this function can only send message to channel
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func printer(c <-chan string) { // <- on left side of chan indicates this function can only recieve message from channel
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second + 1)
	}
}
func main() {
	var c chan string = make(chan string)

	go printer(c)
	go pinger(c)
	go ponger(c)

	var input string
	fmt.Scanln(&input)
}
