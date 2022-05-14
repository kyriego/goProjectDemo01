package main

import (
	"io/ioutil"
	"os"
	"time"
)

func main() {
	ioutil.ReadAll(os.Stdin)
	print(111)
	time.Sleep(20 * time.Second)
}
