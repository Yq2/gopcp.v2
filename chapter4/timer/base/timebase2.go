package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(2 * time.Second)
	fmt.Printf("Present time: %v.\n", time.Now().UTC())
	expirationTime := <-timer.C
	fmt.Printf("Expiration time: %v.\n", expirationTime)
	fmt.Printf("Stop timer: %v.\n", timer.Stop())
	fmt.Println("New timer2 time:", time.Now().UTC())
	timer2 := time.After(2 * time.Second)
	<-timer2
	fmt.Println("Stop timer2 time:", time.Now().UTC())

}
