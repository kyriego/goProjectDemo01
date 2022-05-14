package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var aboard chan struct{} = make(chan struct{})
	var tick = time.NewTicker(1 * time.Second)
	go func() {
		os.Stdin.Read(make([]byte, 1))
		aboard <- struct{}{}
	}()
	for countDown := 10; countDown > 0; countDown-- {
		fmt.Printf("countDown:%d\n", countDown)
		select {
		case <-aboard:
			fmt.Printf("aboard launch!\n")
			return
		case <-tick.C:
		}
	}
	fmt.Printf("launch!\n")
}
