package main

import (
	"fmt"
	"time"
)

func mySleep(duration time.Duration) {
	<-time.After(duration)
}

func main() {
	fmt.Println("Start")
	mySleep(3 * time.Second)
	fmt.Println("Stop")
}
