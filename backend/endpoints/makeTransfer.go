package endpoints

import (
	"WDB/views"
	"fmt"
	"net/http"
)

type TransferData struct {
	value int
	to    string //from będzie obliczane na podstawie tokena :)
}

func (h *Handlers) MakeTransfer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Credentials", "true") // zezwolenie na obsługę ciasteczek
	println("data:\n")
	for _, x := range r.Cookies() {
		fmt.Printf("values %v %v", x.Name, x.Value)
	}
	views.ResponseWithJSON(w, 200, "nice")
}
