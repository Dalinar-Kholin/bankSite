package endpoints

import (
	"WDB/views"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type AceptTransactionRespones struct {
	Link   string `json:"link"`
	From   string `json:"from"`
	To     string `json:"to"`
	Amount int    `json:"amount"`
}

type GetData struct {
	Amount  int    `json:"amount"`
	Reciver string `json:"reciver"`
	Title   string `json:"title"`
}

func (h *Handlers) InitialTransfer(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("accessToken")
	if err != nil {
		views.ResponseWithError(w, 401, "where cookie?")
		return
	}
	var login string
	err = h.DB.QueryRow("select login from users where token = ?", cookie.Value).Scan(&login)

	var data GetData
	bodyReader, _ := io.ReadAll(r.Body)
	err = json.Unmarshal(bodyReader, &data)
	if err != nil {
		fmt.Printf("%v\n", err)
		views.ResponseWithError(w, 400, "bad request?")
		return
	}

	var reciver string
	err = h.DB.QueryRow("select login from users where login = ?", data.Reciver).Scan(&reciver)
	fmt.Printf("--- %v ---\n", data.Reciver)
	if err != nil {
		views.ResponseWithError(w, 401, "bad transaction receiver")
		fmt.Printf("%v\n", err)
		return
	}
	var saldo int
	err = h.DB.QueryRow("select saldo from users where login = ?", data.Reciver).Scan(&saldo)
	if err != nil || saldo < data.Amount {
		views.ResponseWithError(w, 401, "not enought money :(((")
		return
	}

	code := make([]byte, 64)
	if _, err := rand.Read(code); err != nil {
		panic(err)
	}
	h.transactions = append(h.transactions, TransactionData{Title: data.Title, To: data.Reciver, Amount: data.Amount, From: login, Code: hex.EncodeToString(code)})
	views.ResponseWithJSON(w, 200, AceptTransactionRespones{Link: "wysłane do akceptacji"})
	fmt.Printf("%v\n", data.Title)
	//views.ResponseWithJSON(w, 200, AceptTransactionRespones{Link: "https://127.0.0.1:8080/saveTransaction?code=" + hex.EncodeToString(code), From: login, To: data.Reciver, Amount: data.Amount})
}
