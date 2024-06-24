package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	defer func(){
		fmt.Println(time.Since(start))
	}()

	channel := make(chan string)
	go throwingNinjaStar(channel)

	// for message := range channel {
	// 	fmt.Println(message)
	// }

	for {
		message, err := <-channel
		if (!err) {
			break
		}
		fmt.Println(message)
	}
}

func throwingNinjaStar(channel chan string) {
	numRounds := 3
	for i := 0; i < numRounds; i++ {
		score := i
		channel <- fmt.Sprint("Your scored: ", score)
	}
	close(channel)
}