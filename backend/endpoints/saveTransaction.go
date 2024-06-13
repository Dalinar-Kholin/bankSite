package endpoints

import (
	"WDB/views"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ResponseSaveTransfer struct {
	Pass string
}

func (h *Handlers) SaveTransaction(w http.ResponseWriter, r *http.Request) {

	var pass ResponseSaveTransfer
	bodyReader, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(bodyReader, &pass)
	if err != nil {
		views.ResponseWithError(w, 400, "bad request")
		return
	}

	code := r.URL.Query().Get("code")
	ok, data := IsCodeOK(h.transactions, code)
	if !ok {
		views.ResponseWithError(w, 401, "very nie nice")
		return
	}

	var from int
	err = h.DB.QueryRow("select id from users where login = ?", data.From).Scan(&from)
	if err != nil {
		views.ResponseWithError(w, 401, "very nie nice")
		return
	}

	_, isCredsOk := CalculatePassword(h, pass.Pass, data.From)
	if !isCredsOk {
		views.ResponseWithError(w, 401, "bad credentials")
		return
	}

	var to int
	err = h.DB.QueryRow("select id from users where login = ?", data.To).Scan(&to)
	if err != nil {
		views.ResponseWithError(w, 401, "very nie nice")
		return
	}

	res, err := h.DB.Exec("insert into transfers (sender,reciver,value) values (?,?,?);", from, to, data.Amount)
	fmt.Printf("res := %v %v\n", res, err)
	if err != nil {
		views.ResponseWithError(w, 500, "bad server, mea culpa")
		return
	}
	views.ResponseWithJSON(w, 200, "udało sie zrealizowac przelew")

}
