package middlewear

import (
	"WDB/views"
	"fmt"
	"net/http"
	"strconv"
)

func (m *Middlewear) CheckSessionAndToken(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("accessToken")
		if err != nil {
			views.ResponseWithError(w, 401, "where Cookie?")
			return
		}
		var Id int
		var session int64 = 0
		err = m.DB.QueryRow("SELECT id, session FROM users WHERE token=?", cookie.Value).Scan(&Id, &session)
		println(session)
		if err != nil {
			println(err.Error())
			views.ResponseWithError(w, http.StatusUnauthorized, "unauthorize")
			return
		}
		if s, _ := strconv.Atoi(r.Header.Get("Session-Id")); int64(s) != session || session == 0 {
			fmt.Printf("%v %d %d\n", err, s, session)
			views.ResponseWithError(w, http.StatusUnauthorized, "bad session")
			return
		}
		r.Header.Add("id", strconv.Itoa(Id))
		next(w, r)
	}
}
