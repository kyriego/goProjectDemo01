package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"Github.com/kyrieGo/goProjectDemo01/unitconv"
)

//1英尺 = 12英寸 = 0.3048米  1米(m) ~= 3.28084英尺(ft)
func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		f, err := strconv.ParseFloat(input.Text(), 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "tempconv err:%v\n", err)
			continue
		}
		/* 		ct := unitconv.Centigrade(f)
		   		ft := unitconv.CtoF(ct)
		   		kt := unitconv.CtoK(ct)
		   		fmt.Printf("%s = %s = %s\n", ct, ft, kt) */

		meter := unitconv.Meter(f)
		foot := unitconv.MtoF(meter)
		fmt.Printf("%s = %s\n", meter, foot)
	}
}
