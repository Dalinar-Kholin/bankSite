package endpoints

import (
	"WDB/views"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (h *Handlers) GoogleSignIn(w http.ResponseWriter, r *http.Request) {

	var googleData GoogleDataRequest
	bodyReader, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(bodyReader, &googleData)
	fmt.Printf("kinda nice site %v\n", googleData)
	var email string
	err = h.DB.QueryRow("select login from users where email = ?", googleData.Email).Scan(&email)
	if err != sql.ErrNoRows {
		if email != "" {
			views.ResponseWithError(w, 400, "user already exists")
		} else {
			fmt.Printf("%v\n", err)
			views.ResponseWithError(w, 501, "server error")
		}
		return
	}
	_, err = h.DB.Exec("INSERT INTO users (login, email,googleId, saldo) VALUES (?, ?, ?, ?);", googleData.Name, googleData.Email, googleData.GoogleID, 0)
	if err != nil {
		fmt.Printf("err := %v\n", err)
		views.ResponseWithError(w, 431, err.Error())
		return
	}

	views.ResponseWithJSON(w, http.StatusOK, "success")
}

type GoogleDataRequest struct {
	Email      string `json:"email"`
	FamilyName string `json:"familyName"`
	GivenName  string `json:"givenName"`
	GoogleID   string `json:"googleId"`
	Name       string `json:"name"`
}
