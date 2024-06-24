package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type ninja struct {
	name string
}

func main() {
	var sum int64
	fmt.Println(sum)
	atomic.AddInt64(&sum, 1)
	fmt.Println(sum)

	var mu sync.Mutex
	mu.Lock()
	sum = sum + 1
	mu.Unlock()
	fmt.Println(sum)

	var diffSum int64
	fmt.Println(atomic.LoadInt64(&diffSum))
	atomic.StoreInt64(&diffSum, 3)
	atomic.AddInt64(&diffSum, 3)
	fmt.Println(diffSum)

	var av atomic.Value
	wallace := ninja{name: "Wallace"}
	av.Store(wallace)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		w := av.Load().(ninja)
		w.name = "Not Wallace"
		fmt.Println(w)
		av.Store(w)
	}()
	wg.Wait()

	fmt.Println(av.Load().(ninja).name)
}