package unitconv

import "fmt"

type Meter float64
type Foot float64

func (m Meter) String() string {
	return fmt.Sprintf("%.4f m", m)
}

func (f Foot) String() string {
	return fmt.Sprintf("%.4f f", f)
}
