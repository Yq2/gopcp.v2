package main

import (
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	group, errCtx := errgroup.WithContext(ctx)
	var Urls = []string{
		"1",
		"2",
		"3",
		"3",
		"3",
		"4",
	}
	for _, urlVal := range Urls {
		urlVal := urlVal //创建新地址,否则会使用同一个urlVal
		group.Go(func() error {
			err := CheckGoroutineErr(errCtx)
			if err != nil {
				return err
			}
			if urlVal == "2" {
				cancel()
				return err
			}
			fmt.Printf("url: %s\n", urlVal)
			return nil
		})
	}
	if err := group.Wait(); err != nil {
		fmt.Println("err: ", err.Error())
	}
}

func CheckGoroutineErr(errCtx context.Context) error {
	select {
	case <-errCtx.Done():
		return errCtx.Err()
	default:
		return nil
	}
}
