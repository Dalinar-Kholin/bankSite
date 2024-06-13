package endpoints

import (
	"WDB/views"
	"fmt"
	"html"
	"net/http"
	"strconv"
)

type Transfer struct {
	Sender  int
	Reciver int
	Value   int
}

type TransferCalculated struct {
	Sender  string
	Reciver string
	Value   int
}

func (h *Handlers) LoadAllTransfer(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.Header.Get("id"))
	if err != nil {
		fmt.Printf("err := %v", err)
	}
	res, err := h.DB.Query("select sender, reciver, value from transfers where sender = ? || reciver = ?", id, id)
	if err != nil {
		fmt.Printf("err = %v", err)
		views.ResponseWithError(w, 500, "Error querying the database")
		return
	}
	defer res.Close() // Zamknij wynik zapytania na koniec obsługi

	var transfers []Transfer
	var transfersCalculated []TransferCalculated
	for res.Next() {
		var transfer Transfer
		var transferCalculated TransferCalculated
		// zakładając, że tabela 'transfers' ma kolumny 'sender', 'receiver' i 'value'
		if err := res.Scan(&transfer.Sender, &transfer.Reciver, &transfer.Value); err != nil {
			views.ResponseWithError(w, 500, "Error reading data")
			return
		}
		transferCalculated.Value = transfer.Value
		transfersCalculated = append(transfersCalculated, transferCalculated)
		transfers = append(transfers, transfer) // dynamiczne dodawanie do slice
	}

	if err := res.Err(); err != nil {
		views.ResponseWithError(w, 500, "Error after iterating over results")
		return
	}

	nicks, err := h.DB.Query("select distinct users.id,login from users inner join transfers on users.id = transfers.sender || users.id = transfers.reciver where reciver= ? || transfers.sender= ?;", id, id)
	if err != nil {
		fmt.Printf("err = %v", err)
		views.ResponseWithError(w, 501, "bad server")
		return
	}
	mapa := make(map[int]string)
	for nicks.Next() {
		var idFromBase int
		var login string
		// zakładając, że tabela 'transfers' ma kolumny 'sender', 'receiver' i 'value'
		if err := nicks.Scan(&idFromBase, &login); err != nil {
			fmt.Printf("err := %v\n", err)
			views.ResponseWithError(w, 500, "Error reading data")
			return
		}
		fmt.Printf("login %v %v", login, idFromBase)
		mapa[idFromBase] = login
	}
	if err := nicks.Err(); err != nil {
		views.ResponseWithError(w, 500, "wtf err")
		return
	}
	defer nicks.Close()
	for i, x := range transfers {
		transfersCalculated[i].Sender = html.EscapeString(mapa[x.Sender])
		transfersCalculated[i].Reciver = html.EscapeString(mapa[x.Reciver])
	}

	views.ResponseWithJSON(w, 200, transfersCalculated) // odpowiedź JSON z listą transferów
}
