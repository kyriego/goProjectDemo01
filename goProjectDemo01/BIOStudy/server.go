package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func handle(con net.Conn) {
	var bytes []byte = make([]byte, 256)
	for {
		n, err := con.Read(bytes)
		if err != nil {
			if err == io.EOF {
				fmt.Print("数据读取完毕")
			} else {
				fmt.Print("数据读取出错")
			}
			break
		}
		fmt.Printf("%s\n", bytes[:n])
	}

}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatalf("fail to Create net Listener as %s:%v", "localhost:8000", err)
	}
	for {
		c, err2 := l.Accept()
		if err2 != nil {
			log.Printf("getting connection err:%v", err2)
		}
		go handle(c)
		c.Close()
	}
}
