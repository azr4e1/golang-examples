package main

import (
	"fmt"
	"sync"
)

func main() {
	syncMap := sync.Map{}
	syncMap.Store(0, 0)
	syncMap.Store(1, 1)
	syncMap.Store(2, 2)
	syncValue, syncOk := syncMap.Load(0)
	fmt.Println(syncValue, syncOk)
	// for i := 0; i < 10; i++ {
	// 	go func() {
	// 		regularMap[0] = i
	// 	}()
	// }
}
