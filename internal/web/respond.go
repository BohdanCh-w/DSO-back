package web

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Respond(w http.ResponseWriter, status int, v interface{}) error {
	if v == nil {
		w.WriteHeader(status)

		return nil
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("web: write data failed")
	}

	return nil
}
