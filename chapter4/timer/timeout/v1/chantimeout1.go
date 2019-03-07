package main

import (
	"fmt"
	"time"
)

func main() {
	intChan := make(chan int, 1)
	go func() {
		time.Sleep(time.Second)
		intChan <- 1
	}()
	select {
	case e := <-intChan:
		fmt.Printf("Received: %v\n", e)
	case <-time.NewTimer(2 * time.Second).C:
		fmt.Println("Timeout!")
	}
}
