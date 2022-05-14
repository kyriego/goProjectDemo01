package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("c:\\Users\\k7871\\Desktop\\aaa.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "create err:%v\n", err)
	}
	f.Close()
}
