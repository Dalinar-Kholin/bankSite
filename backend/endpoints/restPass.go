package endpoints

import (
	"WDB/checkers"
	"WDB/views"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"gopkg.in/mail.v2"
	"io"
	"net/http"
)

type ResetPassData struct {
	Email string `json:"email"`
	Login string `json:"login"`
}

func sendViaEmail(link string) {
	//wysyłanie linku na email
	m := mail.NewMessage()

	// Zdefiniuj nadawcę, odbiorcę, temat i treść e-maila
	m.SetHeader("From", "osadowskikacper@gmail.com")
	m.SetHeader("To", "gamblebeliever@gmail.com")
	m.SetHeader("Subject", "reset password")
	m.SetBody("text/html", "https://127.0.0.1/newPass/"+"link")

	// Dane serwera SMTP Gmail
	d := mail.NewDialer("smtp.gmail.com", 587, "mail", "pass")

	// Włącz bezpieczne połączenie
	d.StartTLSPolicy = mail.MandatoryStartTLS

	// Wyślij e-mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("mail error := %v", err)
	}

}

func (h *Handlers) ResetPass(w http.ResponseWriter, r *http.Request) {
	var creds ResetPassData
	bodyReader, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(bodyReader, &creds)
	if err != nil {
		views.ResponseWithError(w, 401, "bad data")
		return
	}
	if !checkers.CheckCredsLoginEmail(creds.Login, creds.Email) {
		views.ResponseWithError(w, 405, "bad data")
		return
	}

	var id int
	fmt.Printf("%v %v", creds.Login, creds.Email)
	err = h.DB.QueryRow("select id from users where (login = ? and email = ?)", creds.Login, creds.Email).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			views.ResponseWithError(w, 401, "bad data")
		} else {
			fmt.Println(err)
			views.ResponseWithError(w, 500, "bad server")
		}
		return
	}

	link := make([]byte, 32)
	if _, err := rand.Read(link); err != nil {
		println(err)
	}
	computedLink := hex.EncodeToString(link)
	h.links = append(h.links, ChangePassData{computedLink, id})
	//go sendViaEmail(computedLink)
	views.ResponseWithJSON(w, 200, struct {
		IsOK bool   `json:"isOK"`
		Link string `json:"link"`
	}{
		IsOK: true,
		Link: computedLink,
	})
}
