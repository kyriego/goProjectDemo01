package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	c, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Fprintf(os.Stderr, "dial err:%v\n", err)
	}
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		s := input.Text()
		c.Write([]byte(s))
	}
}
