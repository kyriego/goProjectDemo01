package main

import (
	"fmt"
	"net/http"
	"os"
	"sort"

	"golang.org/x/net/html"
)

//概念们：函数，函数变量的声明，函数变量的创建，匿名函数（获得整个词法环境，可以访问匿名函数外部的变量）
//①定义一个函数 func(a int)int {}
//②函数变量声明（用在函数参数或返回值）  f func(int,float) string
//③创建一个函数变量(直接拿函数名即可)/匿名函数然后调用(相当于是一个函数变量，但是没有类型，类似于java中的匿名内部类（没有类型的类对象）)    func(int)int{blablabla}
//实例：①图的拓扑排序（深度优先搜索）
func add(a int) int {
	return a + 1
}

func Demo01() func() {
	x := "this is the message from client"
	return func() {
		fmt.Printf("%s\n", x)
	}
}

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

//一.先导课程（无向图的拓扑排序，匿名函数的使用）
//我觉得在这里使用“函数变量声明，通过匿名函数定义函数变量，在函数体内通过函数变量调用函数”，起到的是一种将函数和全局变量有序地组织在一起的方式
func toopSort(graph map[string][]string) []string {
	var visited map[string]bool = make(map[string]bool)
	var res []string = make([]string, 0)
	var dfs func([]string)
	dfs = func(items []string) {
		for _, item := range items {
			if !visited[item] {
				visited[item] = true
				dfs(prereqs[item])
				res = append(res, item)
			}
		}
	}
	keys := make([]string, 0, len(graph))
	for k := range graph {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	dfs(keys)
	return res
}

//二.爬虫(1)： 通过一个url得到一堆下一层的url
func Crawl(url string) ([]string, error) {
	fmt.Printf("%s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	root, err1 := html.Parse(resp.Body)
	if err1 != nil {
		return nil, err1
	}
	resp.Body.Close()
	var urls []string = make([]string, 0)
	var ergodicTree func(*html.Node)
	ergodicTree = func(root *html.Node) {
		if root.Type == html.ElementNode && root.Data == "a" {
			for _, attr := range root.Attr {
				if attr.Key == "href" {
					urls = append(urls, attr.Val)
				}
			}
		}
		for c := root.FirstChild; c != nil; c = c.NextSibling {
			ergodicTree(c)
		}
	}
	ergodicTree(root)
	return urls, nil
}

//二.爬虫(2)：广度优先搜索，一串url生成下一串url
func myCrawl(urls []string) { //这个相当于是一个stack
	if len(urls) == 0 {
		return
	}
	visited := make(map[string]bool)
	for len(urls) != 0 {
		tmps := urls //从stack中取出所有元素
		urls = nil
		for _, tmp := range tmps {
			if !visited[tmp] {
				visited[tmp] = true
				s, _ := Crawl(tmp) //处理取出来的元素，得到下一层的元素并加入到stack中
				urls = append(urls, s...)
			}
		}
	}
	return
}

func main() {
	/* 	myCrawl(os.Args[1:]) */
	myCrawl(os.Args[1:])
}
