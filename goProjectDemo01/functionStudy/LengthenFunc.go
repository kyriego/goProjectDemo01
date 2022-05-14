package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

//一.变长函数：
//任务划分：  变长参数定义（strs ...string/ args ...int）  外面如何传递（nums...）   内部如何使用（在内部变长参数是一个slice，但是函数类型不同）
func max(nums ...int) (int, error) {
	if len(nums) < 1 {
		return -1, errors.New("counts of nums must ge 1")
	}
	ans := nums[0]
	for _, num := range nums {
		if num > ans {
			ans = num
		}
	}
	return ans, nil
}

func min(args ...int) (int, error) {
	if len(args) < 1 {
		return -1, errors.New("counts of nums must ge 1")
	}
	res := args[0]
	for _, num := range args {
		if num < res {
			res = num
		}
	}
	return res, nil
}

func myStringsJoin(sep string, strs ...string) string {
	res := ""
	for i, str := range strs {
		res = res + str
		if i != len(strs)-1 {
			res += sep
		}
	}
	return res
}

func check(node *html.Node, name []string) bool {
	if node.Type != html.ElementNode {
		return false
	}
	for _, n := range name {
		if strings.EqualFold(n, node.Data) {
			return true
		}
	}
	return false
}

func ElementByTagName(doc *html.Node, name ...string) []*html.Node {
	res := make([]*html.Node, 0)
	var travel func(node *html.Node)
	travel = func(node *html.Node) {
		if check(node, name) {
			res = append(res, node)
		}
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			travel(c)
		}
	}
	travel(doc)
	return res
}

func main() {
	url := "http://www.baidu.com"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "http Get %s err:%v\n", url, err)
	}
	root, _ := html.Parse(resp.Body)
	resp.Body.Close()
	nodes := ElementByTagName(root, "img", "div")
	fmt.Printf("%d\n", len(nodes))
}
