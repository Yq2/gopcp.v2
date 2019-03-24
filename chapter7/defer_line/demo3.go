package main

import (
	"fmt"
	"time"
)

func main(){
	do()
}

func do() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err: ", err)
		}
	}()

	go da()
	go db()
	time.Sleep(3 * time.Second)
}

func da() {
	panic("panic da")
	for i := 0; i < 10; i++ {
		fmt.Println("i: ", i)
	}
}

func db() {
	//panic("panic db")
	for i := 0; i < 10; i++ {
		fmt.Println("i: ", i)
	}
}
