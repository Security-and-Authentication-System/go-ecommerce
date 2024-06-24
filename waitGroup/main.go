package main

import (
	"fmt"
	"sync"
)

func main() {
	var beeper sync.WaitGroup
	// Names
	names := []string{"Oluwatobi", "Johnson", "Raphael"}
	beeper.Add(len(names))

	for _, name := range names {
		go greeting(name, &beeper)
	}
	beeper.Wait()

	fmt.Println("Mission accomplished")
	fmt.Println("Yessssssssssš")
}

func greeting(name string, beeper *sync.WaitGroup) {
	fmt.Println("Good morning ", name)
	beeper.Done()
}