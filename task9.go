package main

import "fmt"

func main() {
	chOne := make(chan int)
	chTwo := make(chan int)
	digits := [5]int{2, 4, 6, 8, 10}

	go firstChannelWriter(chOne, digits)
	go secondChannelMulti(chOne, chTwo)

	for i := range chTwo {
		fmt.Println(i)
	}
}
func firstChannelWriter(ch chan int, arr [5]int) {
	for i := 0; i < len(arr); i++ {
		ch <- arr[i]
	}
	close(ch)
}
func secondChannelMulti(ch chan int, ch2 chan int) {
	for v := range ch {
		ch2 <- v * 2
	}
	close(ch2)
}
