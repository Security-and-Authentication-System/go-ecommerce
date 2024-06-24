package main

import (
	"fmt"
	"time"
)

// Function that will run concurrently in the main function using goroutine
func say(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond)
	}
}

func main() {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	go say("Hello world!")
	fmt.Println("My second code execution!!!")
	go say("Mr Brine")

	time.Sleep(time.Second)
}