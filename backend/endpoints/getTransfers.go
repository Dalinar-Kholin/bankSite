package endpoints

import (
	"WDB/views"
	"net/http"
)

type Transfer struct {
	Sender  int
	Reciver int
	Value   int
}

func (h *Handlers) GetTransfers(w http.ResponseWriter, r *http.Request) {

	res, err := h.DB.Query("select sender, reciver,value from transfers")
	if err != nil {
		views.ResponseWithError(w, 500, "Error querying the database")
		return
	}
	defer res.Close() // Zamknij wynik zapytania na koniec obsługi

	var transfers []Transfer
	for res.Next() {
		var transfer Transfer
		// zakładając, że tabela 'transfers' ma kolumny 'sender', 'receiver' i 'value'
		if err := res.Scan(&transfer.Sender, &transfer.Reciver, &transfer.Value); err != nil {
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
