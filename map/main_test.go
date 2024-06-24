package main

import (
	"sync"
	"testing"
)

const iterations = 1000

func BenchmarkRegularMap(t *testing.B) {
	var mutex sync.Mutex
	regularMap := make(map[int]int, iterations)
	var wg sync.WaitGroup
	wg.Add(iterations - 1)
	regularMap[0] = 0

	for i := 1; i < iterations; i++ {
		go func() {
			defer wg.Done()
			// j := 1
			for j := 1; j < iterations; j++ {
				go func() {
					mutex.Lock()
					defer mutex.Unlock()
					regularMap[i] = j
				}()
			}
		}()
	}

	wg.Wait()
}

func BenchmarkSyncMark(t *testing.B) {
	var syncMap sync.Map
	var wg sync.WaitGroup
	wg.Add(iterations - 1)
	syncMap.Store(0, 0)

	for i := 1; i < iterations; i++ {
		go func() {
			defer wg.Done()
			// j := 1
			for j := 1; j < iterations; j++ {
				go func() {
					syncMap.Store(i, j)
				}()
			}
		}()
	}

	wg.Wait()
}