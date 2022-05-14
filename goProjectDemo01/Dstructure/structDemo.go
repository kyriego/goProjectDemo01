package main

import (
	"fmt"

	"Github.com/kyrieGo/goProjectDemo01/modules"
)

func main() {
	circle := modules.Circur{
		Point: modules.Point{
			X: 12,
			Y: 25,
		},
		Radius: 125,
	}
	fmt.Printf("%v\n", circle)
}
