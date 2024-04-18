package views

import (
	"encoding/json"
	"net/http"
)

func ResponseWithJSON(w http.ResponseWriter, code int, val interface{}) {
	data, err := json.Marshal(val)
	if err != nil {
		w.WriteHeader(505)
		return
	}
	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(code)
	_, err = w.Write(data)
	if err != nil {

		w.WriteHeader(505)
		return
	}
}

func ResponseWithError(w http.ResponseWriter, code int, val string) {
	type errorResponse struct {
		Error string `json:"error"`
	}
	ResponseWithJSON(w, code, errorResponse{Error: val})
}
