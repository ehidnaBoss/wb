package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	digits_array := [5]int{2, 4, 6, 8, 10}

	var result int

	for _, value := range digits_array {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			val *= val

			mu.Lock()
			result += val
			mu.Unlock()

		}(value)

	}
	wg.Wait()
	fmt.Println(result)
}
