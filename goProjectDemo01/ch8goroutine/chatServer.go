// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 254.
//!+

// Chat is a server that lets clients chat with each other.
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"time"
)

//!+broadcaster
/* type client chan<- string // an outgoing message channel */
type client struct {
	ch   chan string
	name string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				select {
				case cli.ch <- msg:
				default:
					break
				}
			}

		case cli := <-entering:
			if len(clients) > 0 {
				var buffer bytes.Buffer
				buffer.WriteString("users:")
				buffer.WriteByte('[')
				for key, _ := range clients {
					buffer.WriteString(key.name)
					buffer.WriteByte(',')
				}
				buffer.Truncate(buffer.Len() - 1)
				buffer.WriteByte(']')
				cli.ch <- buffer.String()
			} else {
				cli.ch <- "no users!"
			}
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		}
	}
}

//!-broadcaster

//!+handleConn
func handleConn(conn net.Conn) {
	ch := make(chan string, 10) // outgoing client messages
	who := conn.RemoteAddr().String()
	client := client{ch, who}
	done := make(chan struct{})
	go func() {
		for {
			select {
			case <-time.After(20 * time.Second):
				conn.Close()
				return
			case <-done:

			}
		}

	}()
	go clientWriter(conn, client)
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- client

	input := bufio.NewScanner(conn)
	for input.Scan() {
		done <- struct{}{}
		messages <- who + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- client
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, c client) {
	for msg := range c.ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

//!-handleConn

//!+main
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

//!-main
