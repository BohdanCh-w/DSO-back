package usecases

import (
	"github.com/BohdanCh-w/DSO-back/entities"
)

type AnaliticCalculator struct {
	From     float64
	To       float64
	PointNum int
}

func (calc AnaliticCalculator) Calculate() []entities.WawePoint {
	step := (calc.To - calc.From) / float64(calc.PointNum)

	res := make([]entities.WawePoint, 0, calc.PointNum+1)

	for i := calc.From; i <= calc.To; i += step {
		point := entities.WawePoint{X: i, Y: ingergralFunc(i)}
		res = append(res, point)
	}

	return res
}
