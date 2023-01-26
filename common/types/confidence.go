package types

import "fmt"

type Confidence float64

// Applies the math required to obtain the confidence interval value
func (c Confidence) Convert() float64 {
	return float64(c)
}

func (c Confidence) String() string {
	return fmt.Sprintf("%f", c)
}
