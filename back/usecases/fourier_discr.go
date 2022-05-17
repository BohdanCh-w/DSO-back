package usecases

import (
	"math"

	"github.com/BohdanCh-w/DSO-back/entities"
)

type FourierDiscreteCalculator struct {
	From     float64
	To       float64
	PointNum int
	Values   []float64
	points   []entities.WavePoint
}

const shift float64 = 0.206

func (calc FourierDiscreteCalculator) calcA(index int) float64 {
	var sum float64

	for _, p := range calc.points {
		sum += p.Y * math.Cos(p.X*float64(index))
	}

	return sum * 2.0 / float64(len(calc.points))
}

func (calc FourierDiscreteCalculator) calcB(index int) float64 {
	var sum float64

	for _, p := range calc.points {
		sum += p.Y * math.Sin(p.X*float64(index))
	}

	return sum * 2.0 / float64(len(calc.points))
}

func (calc FourierDiscreteCalculator) calcPoints() []entities.WavePoint {
	n := len(calc.Values)
	step := 2 * math.Pi / float64(n)

	points := make([]entities.WavePoint, n)
	for i := 0; i < n; i++ {
		points[i] = entities.WavePoint{X: step * float64(i), Y: calc.Values[i]}
	}

	return points
}

func (calc FourierDiscreteCalculator) Calculate() entities.ResultAnalitics {
	calc.points = calc.calcPoints()

	var (
		n    = len(calc.Values)
		step = (calc.To - calc.From) / float64(calc.PointNum)
		sum  float64
		a0   = calc.calcA(0)

		analitics = entities.ResultAnalitics{
			CoefsA: make([]float64, n),
			CoefsB: make([]float64, n),
			Points: make([]entities.WavePoint, 0, calc.PointNum+1),
		}
	)

	for i := 0; i < n-1; i++ {
		analitics.CoefsA[i] = calc.calcA(i + 1)
		analitics.CoefsB[i] = calc.calcB(i + 1)
	}

	for x := calc.From; x <= calc.To; x += step {
		sum = 0

		for i := 0; i < n-1; i++ {
			sum += analitics.CoefsA[i]*math.Cos(float64(i+1)*(x+shift)) + analitics.CoefsB[i]*math.Sin(float64(i+1)*(x+shift))
		}

		analitics.Points = append(analitics.Points, entities.WavePoint{
			X: x,
			Y: a0/2 + sum,
		})
	}

	return analitics
}
