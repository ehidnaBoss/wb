package main

import (
	"fmt"
	"time"
)

func writeInChannel(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}
func readFromChannel(ch chan int, done chan bool) {
	for {
		select {
		case val := <-ch:
			fmt.Println(val)
		case <-done:
			fmt.Println("Stop")
			return
		}
	}
}
func main() {
	ch := make(chan int)
	done := make(chan bool)

	var n_sec int

	fmt.Println("Введите кол-во сек: ")
	fmt.Scan(&n_sec)

	go writeInChannel(ch)
	go readFromChannel(ch, done)

	time.Sleep(time.Duration(n_sec) * time.Second)

	done <- true
	//close(ch)

}
