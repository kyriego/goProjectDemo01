package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"time"
)

//defer函数： 延迟调用，函数顺序执行到defer语句时，会先计算defer语句中的表达式，然后等到return执行完毕即退出函数时会执行defer的计算结果
//一.用于资源关闭：获取资源成功后通过defer声明释放资源，这样无论如何都会释放资源  defer conn.close()
//二.结束后再执行一个函数：   defer foo()
//三.给函数增加一点入口和出口处理：
//func record(args ...interface{}) func(){}

func record(msg string) func() {
	start := time.Now()
	fmt.Printf("%v: enter %s\n", start, msg)

	return func() {
		end := time.Now()
		fmt.Printf("%v: exit %s(%d ns)\n", end, msg, end.Sub(start))
	}
}

func add(a int) int {
	defer record("add")()
	a++
	time.Sleep(10 * time.Second)
	return a
}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	defer func() {
		err2 := f.Close()
		if err2 != nil {
			if err == nil {
				err = err2
			}
		}
	}()
	n, err = io.Copy(f, resp.Body)
	// Close file, but prefer error from Copy, if any.

	return local, n, err
}

func main() {
	for _, url := range os.Args[1:] {
		local, n, err := fetch(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %s: %v\n", url, err)
			continue
		}
		fmt.Fprintf(os.Stderr, "%s => %s (%d bytes).\n", url, local, n)
	}
}
