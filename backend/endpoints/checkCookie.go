package endpoints

import (
	"WDB/views"
	"database/sql"
	"fmt"
	"net/http"
)

func (h *Handlers) CheckCookie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Credentials", "true") // zezwolenie na obsługę ciasteczek
	cookie, err := r.Cookie("accessToken")
	if err != nil {
		fmt.Printf("error := %v", err)
		views.ResponseWithError(w, 404, "where cookie?")
		return
	}
	var id int
	err = h.DB.QueryRow("select id from users where token=?", cookie.Value).Scan(&id)
	if err == sql.ErrNoRows {
		views.ResponseWithError(w, 401, "bad token go away :(")
		return
	}
	views.ResponseWithJSON(w, 200, CheckCookieResponse{IsValid: true})
}

type CheckCookieResponse struct {
	IsValid bool `json:"isValid"`
}
