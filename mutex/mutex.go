package main

import (
	"fmt"
	"time"
	"sync"
)

var (
	lock sync.Mutex
	rwLock sync.RWMutex
	count int
)

func main() {
	basics()
	readAndWrite()
}

// Function to handle read and write operation
func readAndWrite() {
	go read()
	go write()

	time.Sleep(time.Second * 5)
	fmt.Println("Read and write is done executing!!!")
}

func read() {
	rwLock.RLock()
	defer rwLock.RUnlock()

	fmt.Println("Read locking")
	time.Sleep(1 * time.Second)
	fmt.Println("Reading unlocking")
}

func write() {
	rwLock.Lock()
	defer rwLock.Unlock()

	fmt.Println("Write locking")
	time.Sleep(1 * time.Second)
	fmt.Println("Write unlocking")
}

func basics() {
	iterations := 10000000
	for i := 0; i < iterations; i++ {
		go increment()
	}

	// time.Sleep(time.Second)
	fmt.Println("Result count is: ", count)
}

func increment() {
	lock.Lock()
	count++
	lock.Unlock()
}