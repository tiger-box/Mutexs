package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type ActiveTracker struct {
	counter atomic.Int64
}

func (at *ActiveTracker) Start() {
	at.counter.Add(1)
}

func (at *ActiveTracker) Done() {
	at.counter.Add(-1)
}

func (at *ActiveTracker) Active() int64 {
	return at.counter.Load()
}

func main() {

	traker := ActiveTracker{}

	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for i := 0; i < 1000; i++ {
				traker.Start()
			}

		}()
	}

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				traker.Done()
			}
		}()
	}

	wg.Wait()
	fmt.Println(traker.Active())

}
