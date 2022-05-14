package main

import (
	"Github.com/kyrieGo/goProjectDemo01/unitconv"
	"errors"
	"flag"
	"fmt"
)

type CentigradeFlag struct {
	c unitconv.Centigrade
}

/* func NewCentigradeFlag(c unitconv.Centigrade) *CentigradeFlag {
	centi := new(CentigradeFlag)
	centi.c = c
	return centi
} */
// xxxC  xxx°C   xxxF  xxx°F   xxxK
func (centi *CentigradeFlag) Set(str string) error {
	if len(str) == 0 {
		return errors.New("len str can not be 0!")
	}
	var t string
	var d float64
	fmt.Sscanf(str, "%f%s", &d, &t)
	switch t {
	case "C", "°C":
		centi.c = unitconv.Centigrade(d)
		return nil
	case "F", "°F":
		centi.c = unitconv.FtoC(unitconv.Fahrenheit(d))
		return nil
	case "K":
		centi.c = unitconv.KtoC(unitconv.Kelvins(d))
		return nil
	default:
		return errors.New("temperature pattern error!")
	}
}

func (centi *CentigradeFlag) String() string {
	return centi.c.String()
}

func FlagCentigradeVar(name string, init unitconv.Centigrade, msg string) *unitconv.Centigrade {
	f := CentigradeFlag{init}
	flag.CommandLine.Var(&f, name, msg)
	return &f.c
}

var cen = FlagCentigradeVar("t", 37, "xxxx")

func main() {
	flag.Parse()
	fmt.Printf("%s\n", cen)
}
