package middlewear

import (
	"WDB/views"
	"net/http"
)

func (m *Middlewear) IsAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("accessToken")
		if err != nil {
			views.ResponseWithError(w, 401, "where Cookie?")
			return
		}
		var isAdmin bool
		err = m.DB.QueryRow("select isAdmin from users where token = ?", cookie.Value).Scan(&isAdmin)
		if err != nil || isAdmin == false {
			views.ResponseWithError(w, 400, "only for admins")
			return
		}
		next(w, r)
	}
}
