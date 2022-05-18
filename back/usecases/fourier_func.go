package usecases

import (
	"math"

	"github.com/BohdanCh-w/DSO-back/entities"
	"github.com/phil-mansfield/num"
)

type FourierFuncCalculator struct {
	From       float64
	To         float64
	Iterations int
	PointNum   int
	Func       func(float64) float64
}

func (calc *FourierFuncCalculator) calcA(opts integralOpts) float64 {
	integration := func(x float64) float64 {
		return calc.Func(x) * math.Cos(float64(opts.iter)*x)
	}

	res := num.Integral(integration, opts.from, 1, num.Linear, num.Flat)(opts.to)

	return res / math.Pi
}

func (calc *FourierFuncCalculator) calcB(opts integralOpts) float64 {
	integration := func(x float64) float64 {
		return calc.Func(x) * math.Sin(float64(opts.iter)*x)
	}

	res := num.Integral(integration, opts.from, 1, num.Linear, num.Flat)(opts.to)

	return res / math.Pi
}

func (calc *FourierFuncCalculator) Calculate() entities.ResultAnalitics {
	var (
		step = (calc.To - calc.From) / float64(calc.PointNum)
		sum  float64
		opts = integralOpts{from: -math.Pi, to: math.Pi}
		a0   = calc.calcA(opts)

		analitics = entities.ResultAnalitics{
			CoefsA: make([]float64, calc.Iterations),
			CoefsB: make([]float64, calc.Iterations),
			Points: make([]entities.WavePoint, 0, calc.PointNum+1),
		}
	)

	for i := 0; i < calc.Iterations; i++ {
		opts.iter = i + 1

		analitics.CoefsA[i] = calc.calcA(opts)
		analitics.CoefsB[i] = calc.calcB(opts)
	}

	for x := calc.From; x <= calc.To; x += step {
		sum = 0

		for i := 0; i < calc.Iterations; i++ {
			sum += analitics.CoefsA[i]*math.Cos(float64(i+1)*x) + analitics.CoefsB[i]*math.Sin(float64(i+1)*x)
		}

		analitics.Points = append(analitics.Points, entities.WavePoint{
			X: x,
			Y: a0/2 + sum,
		})
	}

	return analitics
}

func reminder(x float64, y float64) float64 {
	rem := math.Remainder(x, y)
	if rem == 0 {
		return 0
	}
	if rem < 0 {
		return rem + y
	}
	return rem
}

type integralOpts struct {
	iter int
	from float64
	to   float64
}
