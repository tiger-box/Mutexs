package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	counter atomic.Int64
}

func (c *Counter) Increment() {
	c.counter.Add(1)
}

func (c *Counter) Add(n int64) {
	c.counter.Store(n)
}

func (c *Counter) Value() int64 {
	return c.counter.Load()

}

func main() {

	counter := &Counter{}
	const countGorutines = 100
	var wg sync.WaitGroup

	for i := 0; i < countGorutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				counter.Increment()
			}
		}()
	}

	wg.Wait()
	fmt.Println(counter.Value())
	newCount := Counter{}
	newCount.Add(120)
	newCount.Add(200)
	fmt.Println(newCount.Value())
}
