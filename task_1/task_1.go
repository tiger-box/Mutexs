package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	counter int
	mu      sync.Mutex
}

func (c *Counter) Increment() {
	c.mu.Lock()
	c.counter++
	c.mu.Unlock()
}

func (c *Counter) Add(n int) {
	c.mu.Lock()
	c.counter += n
	c.mu.Unlock()
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.counter
}

func main() {

	const countGorutines = 100
	var wg sync.WaitGroup
	counter := &Counter{}

	for i := 0; i < countGorutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				counter.Increment()
				//fmt.Println(counter.counter)
			}
		}()
	}

	wg.Wait()

	newCount := &Counter{}
	newCount.Add(100)
	newCount.Add(1)

	if newCount.Value() == 101 {
		fmt.Println("correct")
	} else {
		fmt.Println("value not correct")
	}
	fmt.Println(newCount.Value())
}
