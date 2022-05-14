package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

//fetch:串行化的fetch请求,将获取的url回传的resp进行打印
//①记录： fetch请求的整体流程： http.Get(url)获取resp  ->    resp的处理 ->         resp.close
func fetch() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: http.Get err: %v\n", err)
			os.Exit(1)
		}
		_, err2 := io.Copy(os.Stdout, resp.Body)
		if err2 != nil {
			fmt.Fprintf(os.Stderr, "fetch: copy err: %v\n", err2)
			os.Exit(1)
		}
		resp.Body.Close()
	}
}

func main() {
	fetch()
}
