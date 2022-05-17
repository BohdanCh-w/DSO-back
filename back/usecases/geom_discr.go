package usecases

import (
	"math"

	"github.com/BohdanCh-w/DSO-back/entities"
)

type GeometricDiscreteCalculator struct {
	From     float64
	To       float64
	PointNum int
	Values   []float64
	points   []entities.WavePoint
}

func (calc GeometricDiscreteCalculator) calcPoints() []entities.WavePoint {
	n := len(calc.Values)
	step := 2 * math.Pi / float64(n)

	points := make([]entities.WavePoint, n+1)
	for i := 0; i < n; i++ {
		points[i] = entities.WavePoint{X: step * float64(i), Y: calc.Values[i]}
	}

	points[n] = entities.WavePoint{X: 2 * math.Pi, Y: calc.Values[0]}

	return points
}

func (calc GeometricDiscreteCalculator) Calculate() []entities.WavePoint {
	var (
		n    = len(calc.Values)
		step = (calc.To - calc.From) / float64(calc.PointNum)

		result = make([]entities.WavePoint, 0, calc.PointNum+1)
	)

	calc.points = calc.calcPoints()

	for x := calc.From; x <= calc.To; x += step {
		var (
			xVal = reminder(x, math.Pi*2)
			yVal float64
		)
		for i := 0; i < n; i++ {
			if xVal == calc.points[i].X {
				yVal = calc.points[i].Y
				break
			}
			if xVal > calc.points[i+1].X {
				continue
			}
			coef := (xVal - calc.points[i].X) / (calc.points[i+1].X - calc.points[i].X)
			yVal = calc.points[i].Y + (calc.points[i+1].Y-calc.points[i].Y)*coef
			break
		}

		result = append(result, entities.WavePoint{X: x, Y: yVal})
	}

	return result
}
