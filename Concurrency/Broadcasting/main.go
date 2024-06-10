package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var ready bool

func main() {
	gettingReadyForMissionCond()
}

func gettingReadyForMissionCond() {
	cond := sync.NewCond(&sync.Mutex{})
	go gettingReadyWithCond(cond)
	workIntervals := 0

	cond.L.Lock()
	defer cond.L.Unlock()

	for !ready {
		workIntervals++
		cond.Wait()
	}

	fmt.Printf("We are now ready! After %d work intervals.\n", workIntervals)
}

func gettingReadyForMission() {
	go gettingReady()
	workIntervals := 0
	for !ready {
		workIntervals++
	}

	fmt.Printf("We are now ready! After %d work intervals.\n", workIntervals)
}

func gettingReady() {
	sleep()
	ready = true
}

func gettingReadyWithCond(cond *sync.Cond) {
	sleep()
	ready = true
	cond.Signal()
}

func sleep() {
	someTime := time.Duration(1+rand.Intn(5)) * time.Second
	time.Sleep(someTime)
}
