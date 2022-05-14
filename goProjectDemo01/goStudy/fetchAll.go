package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

//并行化的fetch请求,得到resp.body后进行时间以及字节数量统计
//①记录：类似barrier的并发处理模式：一个主goroutine要处理众多url请求，对于每个url请求创建出一个goroutine进行处理，每个goroutine处理完毕之后往公共chan string发送一个信息，
//主goroutine接收完所有的从goroutine发送过来的信息之后进行汇总
//③记录： 文件/os.stdout   io.copy   readAll的使用
var cancel chan struct{} = make(chan struct{})

func canceled() bool {
	select {
	case <-cancel:
		return true
	default:
		return false
	}
}

func fetchAll() {
	/* 	start := time.Now() */
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		go fetch1(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	/* 	secs := time.Since(start).Seconds() */
	/* 	fmt.Printf("total time: %.2f s\n", secs) */
}

func fetch1(url string, ch chan string) {
	if canceled() {
		return
	}
	resp, err := http.Get(url)
	if !canceled() {
		close(cancel)
	}
	if err != nil {
		ch <- fmt.Sprintf("get %s err:%v\n", url, err)
	}
	ch <- url
	resp.Body.Close()
}

func main() {
	fetchAll()
}
