package endpoints

import (
	"WDB/views"
	"fmt"
	"net/http"
	"strconv"
)

type Transfer struct {
	Sender  int
	Reciver int
	Value   int
}

func (h *Handlers) GetTransfers(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("accessToken")
	if err != nil {
		views.ResponseWithError(w, 401, "where cookie?")
		return
	}
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	fmt.Printf("id := %v\n")
	res, err := h.DB.Query("select sender from transfers where sender = ? ", id)
	cookie.Value = "asd"
	if err != nil {
		views.ResponseWithError(w, 500, "Error querying the database")
		return
	}
	defer res.Close() // Zamknij wynik zapytania na koniec obsługi

	var transfers []Transfer
	for res.Next() {
		var transfer Transfer
		// zakładając, że tabela 'transfers' ma kolumny 'sender', 'receiver' i 'value'
		if err := res.Scan(&transfer.Sender /*, &transfer.Reciver, &transfer.Value*/); err != nil {
			views.ResponseWithError(w, 500, "Error reading data")
			return
		}

		transfers = append(transfers, transfer) // dynamiczne dodawanie do slice
	}

	if err := res.Err(); err != nil {
		views.ResponseWithError(w, 500, "Error after iterating over results")
		return
	}
	views.ResponseWithJSON(w, 200, transfers) // odpowiedź JSON z listą transferów
}
