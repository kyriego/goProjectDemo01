package main

var cancel chan struct{} = make(chan struct{})

func main() {
	close(cancel)
	close(cancel)
}
