package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//第一阶段：从标准输入中读取所有的文本，统计每一行出现的次数并打印
//记录： ①获取文本（直接就是一个os.stdin）  ②对文本进行一行一行的获取： bufio.newScanner(os.stdin),然后Scan(),Text()读取
func dup1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	showres(counts)
}

//第二阶段：从标准输入和文件列表中进行文本读取，统计每一行出现的次数并打印
//错误记录1：for循环处理一大片文件们，当有一个文件打开出错时，处理错误后要用continue继续循环
//记录： ①获取文本（遍历所有finenames，然后通过os.open(filename)打开文件）
func dup2() {
	index := 1
	counts := make(map[string]int)
	filenames := os.Args[1:]
	if filenames == nil || len(filenames) == 0 {
		scanAndCount(os.Stdin, counts, nil)
	} else {
		for _, filename := range filenames {
			file, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2, fileOpenErr:%v\n", err)
			}
			appear_count := make(map[string]bool)
			flag := scanAndCount(file, counts, appear_count)
			if flag {
				fmt.Printf("RepetitiveFile%d\t%s\n", index, filename)
				index++
			}
			file.Close()
		}
	}
	showres(counts)
}

func scanAndCount(in *os.File, counts map[string]int, appear_count map[string]bool) bool {
	input := bufio.NewScanner(in)
	flag := false
	for input.Scan() {
		text := input.Text()
		counts[text]++
		if appear_count[text] == true {
			flag = true
		}
		appear_count[text] = true
	}
	return flag
}

func showres(counts map[string]int) {
	fmt.Printf("count res:\n")
	for text, count := range counts {
		if count > 1 {
			fmt.Printf("%d\t%s\n", count, text)
		}
	}
}

func dup3() {
	counts := make(map[string]int)
	args := os.Args[1:]
	for _, finename := range args {
		data, err := ioutil.ReadFile(finename)
		if err != nil {

		}
		texts := strings.Split(string(data), "\n")
		for _, text := range texts {
			counts[text]++
		}
	}
	showres(counts)
}

func main() {
	dup2()
}
