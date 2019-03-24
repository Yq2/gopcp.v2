package main

import "fmt"

func main() {
	defer func() {
		if err := recover();err != nil {
			fmt.Println("err: ",err)
		}
	}()

	defer func() {
		// 只有这个panic会触发
		panic("first defer panic")
	}()

	defer func() {
		panic("second defer panic")
	}()

	panic("main body panic")
}


