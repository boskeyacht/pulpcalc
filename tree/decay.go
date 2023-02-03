package tree

import (
	"math"
)

func CalculateDecay(degree int) float64 {
	return math.Log(float64(degree)) / 3
}
