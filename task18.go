package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int
	mu    sync.Mutex
}

func (c *Counter) incrementCounter() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func main() {
	wg := sync.WaitGroup{}

	counter1 := Counter{
		count: 1,
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter1.incrementCounter()
		}()
	}
	wg.Wait()
	fmt.Println(counter1.count)
}
