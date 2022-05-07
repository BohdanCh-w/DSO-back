package usecases

import (
	"math"
)

func ingergralFunc(x float64) float64 {
	reminder := math.Remainder(x, math.Pi*2.0)

	if reminder > 0 {
		return 0
	}
	return 1
}
