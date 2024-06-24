package main

import (
	"fmt"
	"sync"
)

func main() {
	regularMap := make(map[int]interface{})
	syncMap := sync.Map{}

	// put
	regularMap[0] = 10
	regularMap[1] = 11
	regularMap[2] = 22
	syncMap.Store(0, 10)
	syncMap.Store(1, 11)
	syncMap.Store(2, 22)

	// To delete a value in syncMap use LoadAndDelete
	syncValue, loaded := syncMap.LoadAndDelete(2)
	fmt.Println(syncValue, loaded)
	// Getting the value
	regularValue, regularOk := regularMap[0]
	syncValue, syncOk := syncMap.Load(0)
	fmt.Println(regularValue, regularOk)
	fmt.Println(syncValue, syncOk)

	mu := sync.Mutex{}
	mu.Lock()
	// To delete a value in regularMap 
	regularValue = regularMap[2]
	delete(regularMap, 2)
	// fmt.Println(regularMap)
	mu.Unlock()
	fmt.Println(regularValue)

	// Iterating through the regularMap and syncMap
	for key, value := range regularMap {
		fmt.Println(key, value, "||")
	}
	// For syncMap
	syncMap.Range(func(key, value interface{}) bool {
		fmt.Print(key, value, "||")
		return true
	})
	// for i := 0; i < 10; i++ {
	// 	go func() {
	// 		regularMap[0] = i
	// 	}()
	// }
}