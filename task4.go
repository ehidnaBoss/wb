package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {

	mainChan := make(chan int)

	var N int
	fmt.Print("Введите кол-во воркеров: ")
	_, err := fmt.Scan(&N)
	if err != nil {
		fmt.Println("АШИПКА", err)
		return
	}

	wg := sync.WaitGroup{}

	for i := 1; i < N; i++ {
		wg.Add(1)
		go worker(i, mainChan, &wg)
	}
	go func() {
		for i := 1; ; i++ {
			mainChan <- i
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c

	fmt.Println("Получени сигнал завершения...")

	close(mainChan)

	wg.Wait()

	fmt.Println("Воркеры больше не воркают!")
}

func worker(id int, channel <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range channel {
		fmt.Println("worker", id, "started  job", j)
		//time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)

	}
}
