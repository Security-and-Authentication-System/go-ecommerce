package main

import (
	"fmt"
)

func main() {
	x := 2
	if x != 1 {
		fmt.Println("About to Start Panicking")
		panic("Pannicking!!!!!!!!!!!!!!!!!!!!!!!")
	}
	fmt.Println("Code run after x != 1, Hurrayyyyyyyyyy!!!!!!!!!")
}