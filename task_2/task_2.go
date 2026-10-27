package main

import (
	"fmt"
	"sync"
)

type Cashe struct {
	data map[string]string
	mu   sync.RWMutex
}

func (c *Cashe) Set(key, value string) {
	c.mu.Lock()
	c.data[key] = value
	c.mu.Unlock()
}

func (c *Cashe) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.data[key]
	return val, ok
}

func (c *Cashe) Delete(key string) {
	c.mu.Lock()
	if len(c.data) == 0 {
		fmt.Println("eror: len map == 0")
	}
	delete(c.data, key)
	c.mu.Unlock()
}

func Constructor() *Cashe {
	return &Cashe{data: make(map[string]string)}
}

func main() {
	cashe := Constructor()
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			cashe.Set("1", "1")
			cashe.Set("2", "2")
			cashe.Set("3", "3")
			cashe.Set("4", "4")
			cashe.Set("4", "4")
			cashe.Set("6", "6")
			cashe.Set("7", "7")
			cashe.Set("8", "8")
			cashe.Set("9", "9")
			cashe.Set("10", "10")
			cashe.Set("11", "11")
			cashe.Set("12", "12")
			cashe.Set("13", "13")
			cashe.Set("13", "13")
			cashe.Set("14", "14")
			cashe.Set("15", "15")
			cashe.Set("16", "16")
			cashe.Set("17", "17")

			fmt.Println("index", index, " ", "cashe", cashe)
		}(i)
	}

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(cashe.Get("1"))
			fmt.Println(cashe.Get("3"))
			fmt.Println(cashe.Get("5"))
			fmt.Println(cashe.Get("9"))
		}()
	}
	wg.Wait()
}
