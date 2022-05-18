package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/BohdanCh-w/DSO-back/entities"
	"github.com/BohdanCh-w/DSO-back/internal/web"
	"github.com/BohdanCh-w/DSO-back/usecases"
)

func lab3_func1(log *log.Logger, saveLocation string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		web.EnableCors(&w)

		log.Println(r.Host, r.URL.String())

		var req discretePointsRequestWithIter
		if err := req.parse(r); err != nil {
			web.Abort(w, web.NewError(http.StatusBadRequest, err))
			return
		}

		geomCalc := usecases.GeometricDiscreteCalculator{
			From:     req.From,
			To:       req.To,
			PointNum: req.PointNum,
			Values:   req.Values,
		}

		geom := geomCalc.Calculate()

		sqareCalc := usecases.SquareDiscreteCalculator{
			From:     req.From,
			To:       req.To,
			PointNum: req.PointNum,
			Values:   req.Values,
		}

		square, err := sqareCalc.Calculate()
		if err != nil {
			log.Printf("Error calculating square: %v", err)
		}

		coefs := sqareCalc.GetCoefs()

		fourierFunc := func(x float64) float64 {
			return coefs[0] + coefs[1]*x + coefs[2]*x*x
		}

		fourierCalc := usecases.FourierFuncCalculator{
			From:       req.From,
			To:         req.To,
			PointNum:   req.PointNum,
			Iterations: req.Iterations,
			Func:       fourierFunc,
		}

		fouriers := fourierCalc.Calculate()

		response := make([]responsePointTriple, 0, len(fouriers.Points))

		for i := range fouriers.Points {
			response = append(response, responsePointTriple{
				X:  fouriers.Points[i].X,
				Y:  geom[i].Y,
				Yf: fouriers.Points[i].Y,
				Ys: square.Points[i].Y,
			})
		}

		usecases.SaveResult(entities.SaveResult{
			DiffMethodA:      usecases.CalcDifference(geom, fouriers.Points),
			DiffMethodB:      usecases.CalcDifference(geom, square.Points),
			AnaliticsMethodA: fouriers,
			AnaliticsMethodB: square,
		}, saveLocation)

		web.Respond(w, http.StatusOK, response)
	}
}

type discretePointsRequestWithIter struct {
	From       float64
	To         float64
	PointNum   int
	Iterations int
	Values     []float64
}

func (req *discretePointsRequestWithIter) parse(r *http.Request) error {
	query := r.URL.Query()
	var err error

	req.From, err = usecases.ParsePI(query.Get("from"))
	if err != nil {
		return fmt.Errorf("Error reading from parameter: %w", err)
	}

	req.To, err = usecases.ParsePI(query.Get("to"))
	if err != nil {
		return fmt.Errorf("Error reading to parameter: %w", err)
	}

	req.PointNum, err = strconv.Atoi(query.Get("dots"))
	if err != nil {
		return fmt.Errorf("Error reading dots parameter: %w", err)
	}

	req.Iterations, err = strconv.Atoi(query.Get("iterations"))
	if err != nil {
		return fmt.Errorf("Error reading iterations parameter: %w", err)
	}

	vals := strings.Split(query.Get("points"), ",")
	req.Values = make([]float64, len(vals))

	for i, v := range vals {
		req.Values[i], err = strconv.ParseFloat(v, 64)
		if err != nil {
			return fmt.Errorf("Error parsing points parameter: %w", err)
		}
	}

	for key, condition := range map[string]bool{
		"range": req.From < req.To,
		"dots":  req.PointNum > 0,
	} {
		if !condition {
			return fmt.Errorf("%s is not valid", key)
		}
	}

	return nil
}
