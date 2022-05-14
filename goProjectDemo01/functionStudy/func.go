package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"Github.com/kyrieGo/goProjectDemo01/ch6interface"
	"golang.org/x/net/html"
)

//一.访问以root为根的html树，获取所有的href值，并保存在一个slice中（遍历时统计/全局or局部）
//对于局部变量这种思路，我们以前是在开头定义一个全局变量，结尾返回，中间加入当前节点和子节点递归的结果/其实也可以从外面带一个局部变量进来
//记录： 之前学习slice的时候，有一种思路，叫做创建并传入一个slice，返回一个修改后的slice/在刷树的题目的时候，可以用一个slice来接root和子节点执行结果再返回
/* func visited(res []string, root *html.Node) []string {
	if root.Type == html.ElementNode && root.Data == "a" {
		for _, attr := range root.Attr {
			if attr.Key == "href" {
				res = append(res, attr.Val)
			}
		}
	}
	for c := root.FirstChild; c != nil; c = c.NextSibling {
		res = visited(res, c)
	}
	return res
}
*/
//二.访问以root为根的html树，输出一个类似徐html代码结构（遍历时记录状态/全局or局部）
//积累： 以root为当前节点进行遍历时，定义一个状态 depth int / stack []string ,设置并处理当前节点状态，然后传递个下一层节点
//<a>
//	<head>
//</a>
var depth int = 0

func visited1(root *html.Node, depth int) {
	if root.Type == html.ElementNode {
		for i := 0; i < depth; i++ {
			fmt.Print(" ")
		}
		fmt.Printf("<%s>\n", root.Data)
	}
	for c := root.FirstChild; c != nil; c = c.NextSibling {
		visited1(c, depth+1)
	}
	if root.Type == html.ElementNode {
		for i := 0; i < depth; i++ {
			fmt.Print(" ")
		}
		fmt.Printf("</%s>\n", root.Data)
	}
	return
}

//三.广度优先，访问每一个节点，访问到每一个节点，打印从根节点到该节点的路径（遍历时记录状态/全局or局部）
func visited2(root *html.Node, strs []string) {
	if root.Type == html.ElementNode {
		strs = append(strs, root.Data)
		fmt.Println(strs)
	}
	for c := root.FirstChild; c != nil; c = c.NextSibling {
		visited2(c, strs)
	}
}

//四.统计树中所有div标签的个数(遍历时进行统计/全局or局部)  ans:62
//全局统计： 定义一个全局变量，遍历整个树，每次看到一个div标签则对统计值+1
//局部统计①： 在函数体内部定义一个统计变量，接上本节点和子节点们的统计值然后返回
//局部统计②： 传入一个统计变量，统计上本届点，然后传给子节点来统计，子节点统计完后记得更新统计变量，然后返回
var divCount int

func visited3_global(root *html.Node) {
	if root.Type == html.ElementNode && root.Data == "div" {
		divCount++
	}
	for c := root.FirstChild; c != nil; c = c.NextSibling {
		visited3_global(c)
	}
}

func visited3_local(root *html.Node) int {
	divCount := 0
	if root.Type == html.ElementNode && root.Data == "div" {
		divCount++
	}
	for c := root.FirstChild; c != nil; c = c.NextSibling {
		divCount += visited3_local(c)
	}
	return divCount
}

func visited3_local1(root *html.Node, cnt int) int {
	if root.Type == html.ElementNode && root.Data == "div" {
		cnt++
	}
	for c := root.FirstChild; c != nil; c = c.NextSibling {
		cnt = visited3_local1(c, cnt)
	}
	return cnt
}

//五.统计树中各种元素标签的个数
func visited4(root *html.Node, cnt map[string]int) map[string]int {
	if root.Type == html.ElementNode {
		cnt[fmt.Sprintf("<%s>", root.Data)]++
	}
	if root.FirstChild != nil {
		cnt = visited4(root.FirstChild, cnt)
	}
	if root.NextSibling != nil {
		cnt = visited4(root.NextSibling, cnt)
	}
	return cnt
}

//六.输出所有文本节点的内容，但是不包括<script> <style>元素

func visited5(root *html.Node) {
	if root.Type == html.TextNode && root.Data != "script" && root.Data != "style" {
		fmt.Println(root.Data)
	}
	if root.FirstChild != nil {
		visited5(root.FirstChild)
	}
	if root.NextSibling != nil {
		visited5(root.NextSibling)
	}
}

func outText(n *html.Node) {
	if n.Type == html.TextNode {
		fmt.Println(n.Data)
	}
	if c := n.FirstChild; c != nil && c.Data != "style" && c.Data != "script" {
		outText(c)
	}
	if c := n.NextSibling; c != nil && c.Data != "style" && c.Data != "script" {
		outText(c)
	}
}

func expand(s string, sub string, f func(string) string) string {
	if len(s) == 0 || len(sub) == 0 || len(sub) > len(s) {
		return s
	}
	s_b := []byte(s)
	sub_b := []byte(sub)
	i := 0             //    i      index
	for i < len(s_b) { //s_b:  index     index + len(sub_b) - 1
		index := bytes.Index(s_b[i:], sub_b)
		if index == -1 {
			break
		} else {
			tmp := []byte(f(sub))
			copy(s_b[i+index:i+index+len(sub_b)], tmp)
			i = index + len(sub_b)
		}
	}
	return string(s_b)
}

func ElementByID(doc *html.Node, id string) *html.Node {
	return ForEachNode(doc, id)
}

func ForEachNode(root *html.Node, id string) *html.Node {
	if checkNode(root, id) {
		return root
	}
	for c := root.FirstChild; c != nil; c = c.NextSibling {
		return ForEachNode(c, id)
	}
	return nil
}

func checkNode(node *html.Node, id string) bool {
	if node == nil {
		return false
	}
	attrs := node.Attr
	for _, attr := range attrs {
		if attr.Key == "id" && attr.Val == id {
			return true
		}
	}
	return false
}

func visitText(node *html.Node) {
	if node.Type == html.TextNode {
		fmt.Printf("Data:%s\n", node.Data)
		for _, attr := range node.Attr {
			fmt.Printf("%s:%s\n", attr.Key, attr.Val)
		}
		return
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		visitText(c)
	}
}

//
func main() {
	b, _ := ioutil.ReadAll(os.Stdin)
	str := string(b)
	sr := ch6interface.NewMyStringReader(str)
	root, err := html.Parse(sr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "html Parse err:%v\n", err)
		os.Exit(1)
	}
	cnt := make(map[string]int)
	visited4(root, cnt)
	for k, v := range cnt {
		fmt.Printf("%s:%d\n", k, v)
	}
}
