package middlewear

import (
	"net/http"
)

func (m *Middlewear) CheckBodyInput(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r)
	}
}
