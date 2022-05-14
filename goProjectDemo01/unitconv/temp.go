package unitconv

import "fmt"

type Centigrade float64
type Fahrenheit float64
type Kelvins float64

const (
	AbsoluteZeroC Centigrade = -273.15
	BollingPointC Centigrade = 100
	ZeroC         Centigrade = 0
)

func (c Centigrade) String() string {
	return fmt.Sprintf("%g °C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g °F", f)
}

func (k Kelvins) String() string {
	return fmt.Sprintf("%g K", k)
}
