package usecases

import (
	"math"
)

func IngergralFunc(x float64) float64 {
	reminder := math.Remainder(x, math.Pi*2.0)

	if reminder > 0 {
		return 0
	}
	return 1
}

func IngergralFuncFanteak(x float64) float64 {
	return ((math.Pi * math.Pi / 12) - (x * x / 4))
}
