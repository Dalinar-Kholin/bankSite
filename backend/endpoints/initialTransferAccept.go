package endpoints

import (
	"WDB/views"
	"crypto/sha256"
	"encoding/base64"
	"html"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func IsCodeOK(tab []TransactionData, code string) (bool, *TransactionData) {
	for _, x := range tab {
		if x.Code == code {
			return true, &x
		}
	}

	return false, nil
}

func HashAndEncodeBase64(data string) string {
	dataBytes := []byte(data)

	hasher := sha256.New()
	hasher.Write(dataBytes)
	hashBytes := hasher.Sum(nil)

	base64Encoded := base64.StdEncoding.EncodeToString(hashBytes)

	return base64Encoded
}

func HashPage(html string) string {
	start := "<script>"
	end := "</script>"
	startIndex := strings.Index(html, start)
	toHash := html[startIndex+len(start) : strings.Index(html, end)]
	html = strings.Replace(html, "hashPlaceholder", HashAndEncodeBase64(toHash), 1)
	return html
}

func (h *Handlers) InitialTransferAccept(w http.ResponseWriter, r *http.Request) {

	code := r.URL.Query().Get("code")
	ok, data := IsCodeOK(h.transactions, code)
	if !ok {
		views.ResponseWithError(w, 401, "very nie nice")
		return
	}

	filePath := "/home/dalinarkholin/GolandProjects/WDB/backend/endpoints/pages/acceptTransactionPage.html"
	// Odczytanie pliku HTML
	htmlContent, err := os.ReadFile(filePath)
	if err != nil {
		println(err.Error())
		http.Error(w, "Nie można odczytać pliku", 500)
		return
	}
	htmlPage := string(htmlContent)

	// Zastąpienie placeholderów danymi
	htmlPage = strings.Replace(htmlPage, "PlaceholderForFrom", html.EscapeString(data.From), 1)
	htmlPage = strings.Replace(htmlPage, "PlaceholderForTo", html.EscapeString(data.To), 1)
	htmlPage = strings.Replace(htmlPage, "PlaceholderForTitle", html.EscapeString(data.Title), 1)
	htmlPage = strings.Replace(htmlPage, "PlaceholderForAmount", html.EscapeString(strconv.Itoa(data.Amount)), 1)
	//liczenie hasha od scripta aby nie dało się podrobić scripta
	htmlPage = HashPage(htmlPage)

	w.Header().Set("Content-Type", "text/html")

	w.Write([]byte(htmlPage))
}
