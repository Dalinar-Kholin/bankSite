package endpoints

import (
	"WDB/views"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"io"
	"net/http"
	"os"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Money    int    `json:"money"`
	IsAdmin  bool   `json:"isAdmin"`
	jwt.StandardClaims
}

func generateJWT(username, email string, money int, isAdmin bool) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		Email:    email,
		Money:    money,
		IsAdmin:  isAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("secret")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (h *Handlers) GoogleLogIn(w http.ResponseWriter, r *http.Request) {

	var googleData GoogleDataRequest
	bodyReader, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(bodyReader, &googleData)
	if err != nil {
		println(err.Error())
		views.ResponseWithError(w, 502, "server is stupido\n")
		return
	}
	var user string
	err = h.DB.QueryRow("select login from users where email = ?", googleData.Email).Scan(&user)
	if err != nil {
		if err == sql.ErrNoRows {
			_, err = h.DB.Exec("INSERT INTO users (login, email,googleId, saldo) VALUES (?, ?, ?, ?);", googleData.Name, googleData.Email, googleData.GoogleID, 0)
			if err != nil {
				fmt.Printf("err := %v\n", err)
				views.ResponseWithError(w, 431, err.Error())
				return
			}
		} else {
			fmt.Printf("\n?! coto = %v\n", err)
			views.ResponseWithError(w, 501, "server error XD")
			return
		}
	}

	var token string
	var isAdmin bool
	err = h.DB.QueryRow("select token, isAdmin from users where googleId = ?", googleData.GoogleID).Scan(&token, &isAdmin)
	if err == sql.ErrNoRows {
		views.ResponseWithError(w, 403, "there is no user Like that")
		return
	}

	if token == "" {
		newToken := make([]byte, 64)
		if _, err := rand.Read(newToken); err != nil {
			panic(err)
		}
		base64String := base64.StdEncoding.EncodeToString(newToken)
		_, err = h.DB.Exec("update users set token = ? where googleId = ?", base64String, googleData.GoogleID)
		if err != nil {
			views.ResponseWithError(w, 501, "server error co nieXD")
			return
		}
		token = base64String
	}

	session := make([]byte, 4)
	if _, err := rand.Read(session); err != nil {
		panic(err)
	}
	sessionInt := binary.BigEndian.Uint32(session)
	_, err = h.DB.Exec("update users set session = ? where googleId = ?", int64(sessionInt), googleData.GoogleID)
	fmt.Printf("%v \n", int64(sessionInt))
	if err != nil {
		fmt.Printf("%v\n", err)
		views.ResponseWithError(w, 501, "server error co nieXD")
		return
	}

	w.Header().Set("Access-Control-Allow-Credentials", "true") // zezwolenie na obsługę ciasteczek

	myJwt, err := generateJWT(googleData.Name, googleData.Email, 0, isAdmin)
	if err != nil {
		views.ResponseWithError(w, 501, err.Error())
		return
	}

	cookie := &http.Cookie{
		Name:     "accessToken",
		Value:    token,
		Path:     "/",
		SameSite: http.SameSiteNoneMode,
		MaxAge:   3600, // Ważne przez 1 godzinę
		HttpOnly: true, // Ciasteczko dostępne tylko dla serwera
		Secure:   true, // Wymusza używanie HTTPS
	}

	// Ustaw ciasteczko w odpowiedzi
	http.SetCookie(w, cookie)
	views.ResponseWithJSON(w, 200, GoogleResponse{Jwt: myJwt, Session: int64(sessionInt), IsAdmin: isAdmin})
}

type GoogleResponse struct {
	Session int64  `json:"session"`
	Jwt     string `json:"jwt"`
	IsAdmin bool   `json:"isAdmin"`
}
