package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	HaveFunc()
}

func HaveFunc() {
	var Urls = []string{
		"http://wwww.sgfoot.com",
		"http://www.github.com",
	}
	g := &errgroup.Group{}
	for _, urlVal := range Urls{
		urlVal := urlVal //创建新地址,否则会使用同一个urlVal
		g.Go(func() error {
			resp, err := http.Get(urlVal)
			defer resp.Body.Close()
			if err == nil {
				var data []byte
				data, err = ioutil.ReadAll(resp.Body)
				if err != nil {
					return err
				}
				urlRs, err := url.Parse(urlVal)
				if err != nil {
					return err
				}
				fileName := urlRs.Host + ".html"
				fmt.Println(fileName)
				ioutil.WriteFile(fileName, data, 0644)
				return nil
			}
			return err
		})
	}
	if err := g.Wait(); err != nil {
		fmt.Println("err: ", err.Error())
	}
}



