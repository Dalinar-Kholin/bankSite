package endpoints

import (
	"WDB/checkers"
	"WDB/views"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/argon2"
	"io"
	"net/http"
)

//dodać walidacje danych

type UserData struct {
	Password string `json:"pass"`
	Login    string `json:"login"`
	Email    string `json:"email"`
}

func (h *Handlers) AddUser(w http.ResponseWriter, r *http.Request) {
	var userData UserData
	bodyReader, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(bodyReader, &userData)
	if err != nil {
		fmt.Printf("%v\n", err)
		views.ResponseWithError(w, 400, "bad request")
		return
	}
	if !checkers.CheckCredsAll(userData.Password, userData.Email, userData.Login) {
		views.ResponseWithError(w, 400, "bad request")
		return
	}

	var LoginId int
	err = h.DB.QueryRow("select id from users where login = ?", userData.Login).Scan(&LoginId)
	if err != sql.ErrNoRows {
		if LoginId != 0 {
			views.ResponseWithError(w, 400, "user already exists")
		} else {
			fmt.Printf("%v\n", err)
			views.ResponseWithError(w, 501, "server error")
		}
		return
	}

	password := userData.Password
	salt := make([]byte, 64)
	if _, err := rand.Read(salt); err != nil {
		panic(err)
	}

	// Konfiguracja Argon2
	var (
		iterations  uint32 = 3         // Liczba iteracji
		memory      uint32 = 64 * 1024 // Użycie pamięci w KiB
		parallelism uint8  = 4         // Równoległość
		keyLen      uint32 = 32        // Długość klucza
	)

	// Hashowanie hasła
	hash := argon2.Key([]byte(password), salt, iterations, memory, parallelism, keyLen)
	fmt.Printf("argon %v\n sól %v\n", hash, salt)
	computedHash := ComputeHMAC(hex.EncodeToString(hash), h.Pepper)
	_, err = h.DB.Exec("INSERT INTO users (login, passwd, email,salt, saldo) VALUES (?, ?, ?, ?, ?);", userData.Login, computedHash, userData.Email, hex.EncodeToString(salt), 0)
	fmt.Printf("%v\n", err)
	if err != nil {
		views.ResponseWithError(w, 500, "Internal Server Error")
		return
	}

	views.ResponseWithJSON(w, 200, userData)
}
