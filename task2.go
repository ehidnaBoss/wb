package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	digits := [5]int{2, 4, 6, 8, 10}
	for _, value := range digits {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			fmt.Println(val * val)
		}(value)
	}
	wg.Wait()
	//123
}
