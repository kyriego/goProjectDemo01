package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
)

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func main() {
	conn, _ := net.Dial("tcp", "localhost:8001")
	defer conn.Close()
	go mustCopy(os.Stdout, conn)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		str := input.Text()
		conn.Write([]byte(str))
	}
}
