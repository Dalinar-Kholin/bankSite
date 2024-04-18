package main

import (
	"WDB/endpoints"
	"WDB/middlewear"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func makeConToDb() *sql.DB {
	db, err := sql.Open("mysql", "root:j4p13rd0l3@tcp(localhost:3306)/wdb")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {
	// Zmienić ciąg połączenia zgodnie z konfiguracją MariaDB
	godotenv.Load("config.json") // loading .env to let us read it

	DB := makeConToDb()
	defer DB.Close()
	err := DB.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		println("esa")
	}
	handlers := endpoints.Handlers{DB: DB, Pepper: os.Getenv("pepper")}
	midle := middlewear.Middlewear{DB: DB}
	server := http.NewServeMux()
	server.HandleFunc(
		"POST /addUser",
		midle.Cors(handlers.AddUser),
	)
	server.HandleFunc(
		"POST /login",
		midle.Cors(handlers.Login),
	)
	server.HandleFunc(
		"GET /server/passwd",
		midle.Cors(handlers.ResetPassPage),
	)
	server.HandleFunc(
		"POST /server/passwd",
		midle.Cors(handlers.ChangePasswdRequest),
	)
	server.HandleFunc(
		"/transfers",
		midle.Cors(handlers.MakeTransfer))
	server.HandleFunc(
		"/przelewy",
		midle.Cors(midle.CheckToken(handlers.GetTransfers)))

	server.HandleFunc(
		"GET /api/checkCookie",
		midle.Cors(handlers.CheckCookie))
	server.HandleFunc(
		"/resetPassword", //zmienić potem tylko na post, na razie są problemy z CORS
		midle.Cors(handlers.ResetPass))

	/*
		server.HandleFunc(
			"GET /transfers",
			midle.Cors(handlers.GetTransfers))*/
	//http.ListenAndServe(":8080", server)
	err = http.ListenAndServeTLS("127.0.0.1:8080",
		"/home/dalinarkholin/Uczelnia/wstepDoBezpieczenstwa/jebacKlucze/XD/example.com.crt",
		"/home/dalinarkholin/Uczelnia/wstepDoBezpieczenstwa/jebacKlucze/XD/example.com.key",
		server)
	if err != nil {
		log.Fatalf("ListenAndServeTLS error: %v", err)
	}
}
