package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)

// func main () {
// 	gettingReadyForSessionWithCond()	
// }

var ready bool

func gettingReadyForSessionWithCond() {
	cond := sync.NewCond(&sync.Mutex{})
	go gettingReadyWithCond(cond)
	workIntervals := 0

	cond.L.Lock()
	for !ready {
		workIntervals++
		cond.Wait()
	}
	fmt.Printf("We are now ready! After %d work intervals.\n", workIntervals)
	cond.L.Unlock()

}

func gettingReadyWithCond(cond *sync.Cond) {
	sleep()
	ready = true
	cond.Signal()
}

func sleep() {
	rand.Seed(time.Now().UnixNano());
	someTime := time.Duration(1 + rand.Intn(5)) * time.Second
	time.Sleep(someTime)
}