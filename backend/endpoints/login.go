package endpoints

import (
	"WDB/checkers"
	"WDB/views"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"golang.org/x/crypto/argon2"
	"io"
	"net/http"
)

type Creds struct {
	Password string `json:"pass"`
	Login    string `json:"login"`
}

func ComputeHMAC(message, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}

func (h *Handlers) Login(w http.ResponseWriter, r *http.Request) {
	var creds Creds
	bodyReader, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(bodyReader, &creds)
	if err != nil {
		views.ResponseWithError(w, 400, "bad request")
		return
	}
	if !checkers.CheckCredsLoginPass(creds.Login, creds.Password) {
		views.ResponseWithError(w, 400, "bad request")
		return
	}

	var salt string
	err = h.DB.QueryRow("select salt from users where login = ? ", creds.Login).Scan(&salt)
	if err == sql.ErrNoRows {
		views.ResponseWithError(w, 401, "where user?")
		return
	}

	// Konfiguracja Argon2
	var (
		iterations  uint32 = 3         // Liczba iteracji
		memory      uint32 = 64 * 1024 // Użycie pamięci w KiB
		parallelism uint8  = 4         // Równoległość
		keyLen      uint32 = 32        // Długość klucza
	)
	res, _ := hex.DecodeString(salt)
	// Hashowanie hasła
	println(res)
	hash := argon2.Key([]byte(creds.Password), res, iterations, memory, parallelism, keyLen)
	computedHash := ComputeHMAC(hex.EncodeToString(hash), h.Pepper)
	var token string
	err = h.DB.QueryRow("select token from users where login = ? and passwd= ?", creds.Login, computedHash).Scan(&token)
	if err == sql.ErrNoRows {
		views.ResponseWithError(w, 401, "bad credentials")
		return
	}
	if token == "" {
		newToken := make([]byte, 64)
		if _, err := rand.Read(newToken); err != nil {
			panic(err)
		}

		_, err = h.DB.Exec("update users set token = ? where login = ?", hex.EncodeToString(newToken), creds.Login)
		if err != nil {
			views.ResponseWithError(w, 501, "server error")
			return
		}
		token = hex.EncodeToString(newToken)
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
