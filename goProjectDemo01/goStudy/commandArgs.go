package main

import (
	"flag"
	"fmt"
	"os"
)

var n = flag.Int("n", 0, "the number of thread")
var s = flag.Int("s", 0, "ths number of threads")
var f = flag.Bool("f", false, "boolean")

func flagDemo() {
	flag.Parse()
	fmt.Printf("n = %d\n", *n)
	fmt.Printf("s = %d\n", *s)
	fmt.Printf("f = %t\n", *f)
	Args := flag.Args()
	for _, arg := range Args {
		fmt.Println(arg)
	}
	fmt.Println("os.Args:")
	Args = os.Args[1:]
	for _, arg := range Args {
		fmt.Println(arg)
	}
}

func echoDemo() {

}

func main() {
	flagDemo()
}
