package main

import (
	"fmt"
	"time"
)

// Function that will run concurrently in the main function using goroutine
func say(s string, chanl chan bool) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 6000)
	}
	chanl <- true
}

func main() {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	// Make a channel
	channel := make(chan bool, 3)
	go say("Hello world!", channel)
	fmt.Println("My second code execution!!!")
	go say("Good day!", channel)

	fmt.Println("My second code execution!!!ououjihugyyfyfyf")
	channel <- false
	// Getting the message from the channel
	fmt.Println(<-channel)
	fmt.Println(<-channel)

}