package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"
)

type item struct {
	Title string `json:"title"`
	Alt   string `json:"alt"`
}

var wg sync.WaitGroup = sync.WaitGroup{}

func serverInit(MAX int) {
	var infos []item = make([]item, MAX+1)
	for i := 1; i <= MAX; i++ {
		wg.Add(1)
		go queryInfo(i, infos)
	}
	wg.Wait()
	b, err := json.Marshal(infos)
	if err != nil {
		fmt.Fprintf(os.Stderr, "infos marshal err:%v\n", err)
	}
	f, err2 := os.Create("infos.txt")
	if err2 != nil {
		fmt.Fprintf(os.Stderr, "file create err:%v\n", err2)
	}
	_, err3 := f.Write(b)
	if err3 != nil {
		fmt.Fprintf(os.Stderr, "file write err:%v\n", err2)
	}
}

func queryInfo(i int, infos []item) {
	defer wg.Done()
	IssuesURL := fmt.Sprintf("https://xkcd.com/%d/info.0.json", i)
	resp, err := http.Get(IssuesURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "httpGet err: %s\n", IssuesURL)
		return
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		fmt.Fprintf(os.Stderr, "search query failed: %s\n", resp.Status)
		return
	}
	var result item
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return
	}
	infos[i] = result
}

func query() {
	args := os.Args[1:]
	if len(args) <= 0 || len(args) > 2 || args[0] > args[1] {
		fmt.Fprint(os.Stderr, "args err!")
		return
	}
	f, err := os.Open("infos.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "file open err:%v\n", err)
	}
	b2, err2 := ioutil.ReadAll(f)
	if err2 != nil {
		fmt.Fprintf(os.Stderr, "file read err:%v\n", err2)
	}
	var infos []item
	json.Unmarshal(b2, &infos)
	a, _ := strconv.Atoi(args[0])
	b, _ := strconv.Atoi(args[1])
	l, r := maxmin(a, b)
	for i := l; i <= r; i++ {
		fmt.Printf("%d:\t Title:%s|Alt:%s\n", i, infos[i].Title, infos[i].Alt)
	}
}

func maxmin(a int, b int) (int, int) {
	if a < b {
		a, b = b, a
	}
	return b, a
}

/* func init() {
	serverInit(500)
} */

func main() { // 1  2  (a < b)
	query()
}
