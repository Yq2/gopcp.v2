package main

import "fmt"

// 这个捕获会失败
defer recover()

defer fmt.Println("这个捕获也会失败")

defer func() {
	func() {
		// 这个间接使用的recover无效
		println("defer inner ")
		recover()
	}()
}()

defer func() {
	// 有效
	fmt.Println("defer inner")
	recover()
}()

func except() {
	recover()
}

func test(){
	defer except
}

