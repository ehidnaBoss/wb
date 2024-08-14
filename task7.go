package main

import (
	"fmt"
	"sync"
)

func main() {
	m := make(map[int]int)
	var mu = sync.Mutex{}
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(k int) {
			defer wg.Done()

			mu.Lock()
			m[k] = k * k
			mu.Unlock()
		}(i)
	}
	wg.Wait()

	for k, v := range m {
		fmt.Println(k, v)
	}
}
