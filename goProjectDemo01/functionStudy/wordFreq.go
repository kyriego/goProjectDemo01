package main

import (
	"bufio"
	"fmt"
	"os"
)

func wordFreq(filename string, cnt map[string]int) (map[string]int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	input := bufio.NewScanner(f)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		s := input.Text()
		cnt[s]++
	}
	return cnt, nil
}

func main() {
	cnt, _ := wordFreq("article.txt", make(map[string]int))
	for k, v := range cnt {
		fmt.Printf("%s:%d\n", k, v)
	}
}
