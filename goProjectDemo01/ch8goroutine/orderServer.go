package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
)

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		var b []byte = make([]byte, 512)
		n, err := conn.Read(b)
		if err == io.EOF {
			log.Printf("read bytes from %s EOF:%v", conn.RemoteAddr(), err)
			break
		} else if err != nil && err != io.EOF {
			log.Printf("read bytes from %s err:%v", conn.RemoteAddr(), err)
			break
		}
		b = b[0:n]
		b = bytes.TrimSpace(b)
		if bytes.HasPrefix(b, []byte("cd")) {
			b = bytes.TrimSpace(b[2:])
			if len(b) == 0 {
				fmt.Printf("cd args err!\n")
				conn.Write([]byte("cd args err!\n"))
				continue
			}
			fmt.Printf("cd %s success!\n", b)
			conn.Write([]byte(fmt.Sprintf("cd %s success!\n", b)))
		} else if bytes.HasPrefix(b, []byte("ls")) {
			b = bytes.TrimSpace(b[2:])
			if len(b) == 0 {
				fmt.Printf("ls args err!\n")
				conn.Write([]byte("ls args err!\n"))
				continue
			}
			fmt.Printf("ls %s success!\n", b)
			conn.Write([]byte(fmt.Sprintf("ls %s success!\n", b)))
		} else if bytes.HasPrefix(b, []byte("get")) {
			b = bytes.TrimSpace(b[3:])
			if len(b) == 0 {
				fmt.Printf("get args err!\n")
				conn.Write([]byte("get args err!\n"))
				continue
			}
			fmt.Printf("get %s success!\n", b)
			conn.Write([]byte(fmt.Sprintf("get %s success!\n", b)))
		} else if bytes.HasPrefix(b, []byte("close")) {
			b = bytes.TrimSpace(b[5:])
			if len(b) == 0 {
				fmt.Printf("close args err!\n")
				conn.Write([]byte("close args err!\n"))
				continue
			}
			fmt.Printf("close %s success!\n", b)
			conn.Write([]byte(fmt.Sprintf("close %s success!\n", b)))
		} else {
			fmt.Printf("unexpected command!\n")
			conn.Write([]byte("unexpected command!\n"))
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8001")
	if err != nil {
		log.Fatalf("net listen err:%v", err)
	}
	for {
		conn, err2 := listener.Accept()
		if err2 != nil {
			log.Printf("accept conn from %s err:%v", conn.RemoteAddr(), err2)
			continue
		}
		go handleConn(conn)
	}
}
