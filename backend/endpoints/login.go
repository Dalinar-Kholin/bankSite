package endpoints

import (
	"WDB/views"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/Dalinar-Kholin/sqlLoger"
	"golang.org/x/crypto/argon2"
	"io"
	"net/http"
)

type Creds struct {
	Password string `json:"pass"`
	Login    string `json:"login"`
}

func ComputeHMAC(message, key string) (res string) {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(message))
	res = hex.EncodeToString(h.Sum(nil))
	return
}

func CalculatePassword(h *Handlers, password, login string) (string, bool) {
	var salt string

	// err := h.DB.QueryRow("select salt from users where login = ? ", login).Scan(&salt) // na czas sql injection
	err := sqlLoger.QueryRow("select salt from users where login = ?", login).Scan(&salt)
	if err == sql.ErrNoRows {
		return "server is stupid", false
	}

	// Konfiguracja Argon2
	var (
		iterations  uint32 = 3         // Liczba iteracji
		memory      uint32 = 64 * 1024 // Użycie pamięci w KiB
		parallelism uint8  = 4         // Równoległość
		keyLen      uint32 = 32        // Długość klucza
	)
	res, _ := hex.DecodeString(salt)
	hash := argon2.Key([]byte(password), res, iterations, memory, parallelism, keyLen)
	computedHash := ComputeHMAC(hex.EncodeToString(hash), h.Pepper)
	var token string
	err = sqlLoger.QueryRow("select token from users where login = ? and passwd= ?", login, computedHash).Scan(&token)
	//err = h.DB.QueryRow("select token from users where login = ? and passwd= ?", login, computedHash).Scan(&token)
	if err == sql.ErrNoRows {
		return "bad creds", false
	}
	return token, true
}

func (h *Handlers) Login(w http.ResponseWriter, r *http.Request) {
	var creds Creds
	bodyReader, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(bodyReader, &creds)
	if err != nil {
		views.ResponseWithError(w, 400, "bad request")
		return
	}
	/*	if !checkers.CheckCredsLoginPass(creds.Login, creds.Password) {
		views.ResponseWithError(w, 405, "bad request")
		return
	} aby udaly sie ataki sql injection trzeba bylo to wyciac */

	token, isCredsOk := CalculatePassword(h, creds.Password, creds.Login)
	if !isCredsOk {
		views.ResponseWithError(w, 400, token)
		return
	}

	if token == "" {
		newToken := make([]byte, 64)
		if _, err := rand.Read(newToken); err != nil {
			panic(err)
		}
		base64String := base64.StdEncoding.EncodeToString(newToken)
		_, err = h.DB.Exec("update users set token = ? where login = ?", base64String, creds.Login)

		if err != nil {
			views.ResponseWithError(w, 500, "server error")
			return
		}
		token = base64String
	}
	session := make([]byte, 4)
	if _, err := rand.Read(session); err != nil {
		panic(err)
	}
	sessionInt := binary.BigEndian.Uint32(session)
	_, err = h.DB.Exec("update users set session = ? where login = ?", int64(sessionInt), creds.Login)
	fmt.Printf("%v \n", int64(sessionInt))
	if err != nil {
		fmt.Printf("%v\n", err)
		views.ResponseWithError(w, 501, "server error co nieXD")
		return
	}

	w.Header().Set("Access-Control-Allow-Credentials", "true") // zezwolenie na obsługę ciasteczek

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
	views.ResponseWithJSON(w, 200, "succesfull Login")
}
