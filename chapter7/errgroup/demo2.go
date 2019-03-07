package main

import (
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
)

func main() {
	//创建一个错误组
	eg := &errgroup.Group{}
	//将需要执行的方法添加到Go函数里,实现并发
	eg.Go(func() error {
		return PrintNumber("go", 1)
	})
	eg.Go(func() error {
		return PrintNumber("php", 999)
	})
	//扑捉错误,注: 协程之间其中一个出错不会影响另一个协程的执行结果
	if err := eg.Wait(); err != nil {
		fmt.Println(err.Error())
	}
}

//一个打印函数, 传个标识
func PrintNumber(flag string, errNumber int) error {
	i := 1
	for i <= 1000 {
		if i == errNumber {
			err := errors.New("my is error, flag: " + flag)
			return err
		}
		fmt.Printf("flag:%s, number:%d\n", flag, i)
		i++
	}
	return nil
}
