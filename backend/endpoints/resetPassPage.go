package endpoints

import (
	"WDB/views"
	"fmt"
	"net/http"
	"os"
)

func Contains(tab []ChangePassData, link string) (bool, int) {
	for i := 0; i < len(tab); i++ {
		if tab[i].link == link {
			return true, tab[i].id
		}
	}
	return false, -1
}

func (h *Handlers) ResetPassPage(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%v %v\n", r.URL.Query().Get("code"), h.links)
	isOK, _ := Contains(h.links, r.URL.Query().Get("code"))
	if !isOK {
		views.ResponseWithError(w, 401, "nieładnie :(")
		return
	}

	filePath := "/home/dalinarkholin/GolandProjects/WDB/backend/endpoints/page.html" // Zmień "twoja_strona.html" na nazwę Twojego pliku HTML
	// Odczytanie pliku HTML
	htmlContent, err := os.ReadFile(filePath)
	if err != nil {
		println(err.Error())
		http.Error(w, "Nie można odczytać pliku", 500)
		return
	}

	// Ustawienie typu treści odpowiedzi na HTML
	w.Header().Set("Content-Type", "text/html")
	w.Write(htmlContent)
}
