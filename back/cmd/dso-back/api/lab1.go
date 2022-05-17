package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/BohdanCh-w/DSO-back/entities"
	"github.com/BohdanCh-w/DSO-back/internal/web"
	"github.com/BohdanCh-w/DSO-back/usecases"
)

func lab1_func1(log *log.Logger, saveLocation string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		web.EnableCors(&w)

		log.Println(r.Host, r.URL.String())

		var req fourierFuncRequest
		if err := req.parse(r); err != nil {
			web.Abort(w, web.NewError(http.StatusBadRequest, err))
			return
		}

		analitic := usecases.AnaliticCalculator{
			From:     req.From,
			To:       req.To,
			PointNum: req.PointNum,
		}
		analitics := analitic.Calculate()

		fourier := usecases.FourierFuncCalculator{
			From:       req.From,
			To:         req.To,
			Iterations: req.Iterations,
			PointNum:   req.PointNum,
		}
		fouriers := fourier.Calculate()

		response := make([]responsePointDouble, 0, len(analitics))

		for i := range analitics {
			response = append(response, responsePointDouble{
				X:  analitics[i].X,
				Y:  analitics[i].Y,
				Yf: fouriers.Points[i].Y,
			})
		}

		usecases.SaveResult(entities.SaveResult{
			DiffMethodA:      usecases.CalcDifference(analitics, fouriers.Points),
			AnaliticsMethodB: fouriers,
		}, saveLocation)

		web.Respond(w, http.StatusOK, response)
	}
}

type responsePointDouble struct {
	X  float64 `json:"x"`
	Y  float64 `json:"y"`
	Yf float64 `json:"yf"`
}

type fourierFuncRequest struct {
	From       float64
	To         float64
	Iterations int
	PointNum   int
}

func (req *fourierFuncRequest) parse(r *http.Request) error {
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

	req.Iterations, err = strconv.Atoi(query.Get("iterations"))
	if err != nil {
		return fmt.Errorf("Error reading iterations parameter: %w", err)
	}

	req.PointNum, err = strconv.Atoi(query.Get("dots"))
	if err != nil {
		return fmt.Errorf("Error reading dots parameter: %w", err)
	}

	if req.From >= req.To {
		return fmt.Errorf("Ending point must be greater than stating")
	}

	for key, value := range map[string]float64{
		"iterations": float64(req.Iterations),
		"dots":       float64(req.PointNum),
	} {
		if value <= 0 {
			return fmt.Errorf("%s must be greater than 0", key)
		}
	}

	return nil
}
