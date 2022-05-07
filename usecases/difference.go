package usecases

import (
	"math"

	"github.com/BohdanCh-w/DSO-back/entities"
)

func CalcDifference(a, b []entities.WawePoint) (res float64) {
	for i := 0; i < len(a); i++ {
		res += math.Abs(a[i].Y) - math.Abs(b[i].Y)
	}

	return
}
