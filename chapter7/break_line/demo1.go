package main

import "fmt"

func main() {

	fmt.Println("break 语句规则展示")
L1:
	for i := 0; ; i++ {
		for j := 0; ; j++ {
			if i > 5 {
				break L1
			}
			if j > 10 {
				break
			}
			fmt.Printf("i=%d,j=%d\n", i, j)
		}
	}

	fmt.Println("contine 语句规则展示")
L2:
	for i := 0; ; i++ {
		for j := 0; ; j++ {
			fmt.Printf("i=%d,j=%d\n", i, j)
			if i >= 5 {
				continue L2
			}
			if j > 10 {
				continue
			}
		}
	}
	fmt.Println("展示结束")
}
