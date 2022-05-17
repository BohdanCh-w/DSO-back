package usecases

import (
	"fmt"
	"math"

	"github.com/BohdanCh-w/DSO-back/entities"
	"github.com/shopspring/decimal"
)

type SquareDiscreteCalculator struct {
	From     float64
	To       float64
	PointNum int
	Values   []float64
	points   []entities.WavePoint
}

func (calc SquareDiscreteCalculator) calcPoints() []entities.WavePoint {
	n := len(calc.Values)
	step := 2 * math.Pi / float64(n)

	points := make([]entities.WavePoint, n+1)
	for i := 0; i < n; i++ {
		points[i] = entities.WavePoint{X: step * float64(i), Y: calc.Values[i]}
	}

	points[n] = entities.WavePoint{X: 2 * math.Pi, Y: calc.Values[0]}

	return points
}

func (calc SquareDiscreteCalculator) buildMatrix() [][]float64 {
	var (
		n        = 3
		matrix   = make([][]float64, n)
		cacheSum = make(map[int]float64)
	)

	for i := range matrix {
		matrix[i] = make([]float64, n+1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if val, ok := cacheSum[i+j]; ok {
				matrix[i][j] = val
				continue
			}

			var sumX float64
			for _, p := range calc.points {
				sumX += math.Pow(p.X, float64(i+j))
			}
			matrix[i][j] = sumX
			cacheSum[i+j] = sumX
		}

		var sumY float64
		for _, p := range calc.points {
			sumY += p.Y * math.Pow(p.X, float64(i))
		}
		matrix[i][n] = sumY
	}

	return matrix
}

func (calc SquareDiscreteCalculator) Gaus(matrix [][]float64) bool {
	var (
		ok bool
		n  = len(matrix)
	)

	for i := 0; i < n; i++ {
		if matrix[i][i] == 0 {
			var c = 1
			for (i+c) < n && matrix[i+c][i] == 0 {
				c++
			}

			if (i + c) == n {
				ok = true
				break
			}

			for j, k := i, 0; k <= n; k++ {
				matrix[j][k], matrix[j+c][k] = matrix[j+c][k], matrix[j][k]
			}
		}

		for j := 0; j < n; j++ {
			if i != j {
				p := matrix[j][i] / matrix[i][i]

				for k := 0; k <= n; k++ {
					matrix[j][k] -= matrix[i][k] * p
				}
			}
		}
	}

	return ok
}

func (calc SquareDiscreteCalculator) checkConsistency(matrix [][]float64) int {
	var n = len(matrix)

	for i := 0; i < n; i++ {
		var sum float64
		var j int
		for j = 0; j < n; j++ {
			sum += matrix[i][j]

		}
		if decimal.NewFromFloat(sum).Equal(decimal.NewFromFloat(matrix[i][j])) {
			return 2
		}
	}

	return 3
}

func (calc SquareDiscreteCalculator) calcCoef(matrix [][]float64, power int) ([]float64, error) {
	if power == 2 {
		return nil, fmt.Errorf("Infinite Solutions possible")
	}
	if power == 3 {
		return nil, fmt.Errorf("No solution exist")
	}

	var (
		n      = len(matrix)
		result = make([]float64, n)
	)

	for i := 0; i < n; i++ {
		result[i] = matrix[i][n] / matrix[i][i]
	}

	return result, nil
}

func (calc SquareDiscreteCalculator) Calculate() (entities.ResultAnalitics, error) {
	calc.points = calc.calcPoints()
	var (
		n    = len(calc.Values)
		step = (calc.To - calc.From) / float64(calc.PointNum)

		analitics = entities.ResultAnalitics{
			CoefsA: make([]float64, n),
			CoefsB: make([]float64, n),
			Points: make([]entities.WavePoint, 0, calc.PointNum+1),
		}

		matrix = calc.buildMatrix()
	)

	var power int
	if ok := calc.Gaus(matrix); ok {
		power = calc.checkConsistency(matrix)
	}

	coefs, err := calc.calcCoef(matrix, power)
	if err != nil {
		return analitics, err
	}

	for x := calc.From; x <= calc.To; x += step {
		var y float64

		for i, c := range coefs {
			y += c * math.Pow(x, float64(i))
		}

		analitics.Points = append(analitics.Points, entities.WavePoint{
			X: x,
			Y: y,
		})
	}

	return analitics, nil
}
