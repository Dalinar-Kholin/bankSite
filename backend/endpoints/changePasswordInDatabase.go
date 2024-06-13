package endpoints

import (
	"WDB/checkers"
	"WDB/views"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/argon2"
	"io"
	"net/http"
)

type Pass struct {
	Pass   string `json:"pass"`
	PassCp string `json:"passCp"`
}

func (h *Handlers) ChangePasswordInDatabase(w http.ResponseWriter, r *http.Request) {
	isOK, id := Contains(h.links, r.URL.Query().Get("code"))
	if !isOK {
		views.ResponseWithError(w, 401, "bad link")
		return
	}
	var Passes Pass
	bodyReader, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(bodyReader, &Passes)

	if !checkers.ValidatePassword(Passes.PassCp) && !checkers.ValidatePassword(Passes.Pass) {
		views.ResponseWithError(w, 403, "bad data")
		return
	}

	if err != nil {
		fmt.Println(err)
		views.ResponseWithError(w, 402, "bad data serio>")
		return
	}
	if Passes.Pass != Passes.PassCp {
		views.ResponseWithError(w, 405, "password npt matching")
		return
	}
	var salt string
	err = h.DB.QueryRow("select salt from users where id=?", id).Scan(&salt)
	if err != nil {
		if err == sql.ErrNoRows {
			views.ResponseWithError(w, 401, "wtf")
		} else {
			views.ResponseWithError(w, 403, "wtf")
		}
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
	fmt.Printf("haslo := %v\n", Passes.Pass)
	hash := argon2.Key([]byte(Passes.Pass), res, iterations, memory, parallelism, keyLen)
	computedHash := ComputeHMAC(hex.EncodeToString(hash), h.Pepper)
	_, err = h.DB.Exec("update users set passwd = ? where id = ?", computedHash, id)

	fmt.Printf("chaslo do zmiany")
	views.ResponseWithJSON(w, 200, "allOK")
}
