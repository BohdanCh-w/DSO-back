package web

import (
	"net/http"
)

func Abort(w http.ResponseWriter, weberr *Error) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(weberr.Code)

	w.Write([]byte(`{"error": "` + weberr.Error() + `"}`))

	return nil
}
