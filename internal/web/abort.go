package web

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Abort(w http.ResponseWriter, weberr *Error) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(weberr.Code)

	if err := json.NewEncoder(w).Encode(weberr); err != nil {
		return fmt.Errorf("web: write data failed")
	}

	return nil
}
