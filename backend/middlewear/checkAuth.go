package middlewear

import (
	"WDB/views"
	"net/http"
	"strconv"
)

func (m *Middlewear) CheckToken(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("accessToken")
		if err != nil {
			views.ResponseWithError(w, 401, "where Cookie?")
			return
		}
		var Id int
		err = m.DB.QueryRow("SELECT id FROM users WHERE token=?", cookie.Value).Scan(&Id)
		if err != nil {
			println(err.Error())
			views.ResponseWithError(w, http.StatusUnauthorized, "unauthorize")
			return
		}
		r.Header.Add("id", strconv.Itoa(Id))
		next(w, r)
	}
}
