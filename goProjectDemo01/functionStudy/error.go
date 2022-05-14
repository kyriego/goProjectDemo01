package main

func foo2(x int) {
	print(0 / x)
}

func foo1() {
	foo2(0)
}
func main() {
	foo1()
}
