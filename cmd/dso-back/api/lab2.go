package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/BohdanCh-w/DSO-back/internal/web"
	"github.com/BohdanCh-w/DSO-back/usecases"
)

func lab2_func1(log *log.Logger, saveLocation string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		web.EnableCors(&w)

		log.Println(r.Host, r.URL.String())

		var req fourierDiscreteRequest
		if err := req.parse(r); err != nil {
			web.Abort(w, web.NewError(http.StatusBadRequest, err))
			return
		}

		web.Respond(w, http.StatusOK, req)
	}
}

type fourierDiscreteRequest struct {
	From     float64
	To       float64
	PointNum int
	Points   []float64
}

func (req *fourierDiscreteRequest) parse(r *http.Request) error {
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

	vals := strings.Split(query.Get("points"), ",")
	req.Points = make([]float64, len(vals))

	for i, v := range vals {
		req.Points[i], err = strconv.ParseFloat(v, 64)
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
