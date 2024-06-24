package main

import (
	"fmt"
	"sync"
)

func main() {
	gettingReadyForSessionWithCond()	
	broadCastStartOfMission()
}

func broadCastStartOfMission() {
	beeper := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup
	wg.Add(3)
	standByForMission(func(){
		fmt.Println("Ninja 1 starting mission.")
		wg.Done()
	}, beeper)
	standByForMission(func() {
		fmt.Println("Ninja 2 starting mission.")
		wg.Done()
	}, beeper)
	standByForMission(func() {
		fmt.Println("Ninja 3 starting mission")
		wg.Done()
	}, beeper)
	beeper.Broadcast()
	wg.Wait()
	fmt.Println("All Ninjas have started their missions")
}

func standByForMission(fn func(), beeper *sync.Cond) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		wg.Done()
		beeper.L.Lock()
		defer beeper.L.Unlock()
		beeper.Wait()
		fn()
	}()
	wg.Wait()
}