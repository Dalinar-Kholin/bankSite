package main

import (
	"WDB/endpoints"
	"WDB/middlewear"
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"fmt"
	"github.com/Dalinar-Kholin/sqlLoger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func makeConToDb(dbConnect string) *sql.DB {
	db, err := sql.Open("mysql", dbConnect)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {
	// Zmienić ciąg połączenia zgodnie z konfiguracją MariaDB

	err := godotenv.Load(".env")
	if err != nil {
		println(err.Error())
		return
	} // loading .env to let us read it

	if e := sqlLoger.SetUpLogger("sql-log", os.Getenv("dbConnect")); e != nil { // ICOM: setupowanie biblioteki sqlLoger aby poprawnie działała w googleLogin
		println(e.Error())
	}

	DB := makeConToDb(os.Getenv("dbConnect"))
	defer DB.Close()
	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		println("esa")
	}
	handlers := endpoints.Handlers{DB: DB, Pepper: os.Getenv("pepper")}
	midle := middlewear.Middlewear{DB: DB}
	server := http.NewServeMux()
	var isAdmin bool
	DB.QueryRow("select isAdmin from users where id = 33 ").Scan(&isAdmin)
	fmt.Printf("isAdmin = %v\n", isAdmin)
	// register new user
	server.HandleFunc(
		"POST /api/registerNewUser",
		midle.Cors(handlers.RegisterNewUser),
	)
	// login
	server.HandleFunc(
		"POST /login",
		midle.Cors(midle.CheckBodyInput(handlers.Login)),
	)

	// reset password
	server.HandleFunc(
		"/server/password",
		midle.Cors(handlers.ResetPassAccept),
	)
	server.HandleFunc(
		"/server/passwd",
		midle.Cors(handlers.ChangePasswordInDatabase),
	)

	server.HandleFunc(
		"/google",
		midle.Cors(handlers.GoogleLogIn),
	)

	server.HandleFunc(
		"/transfers",
		midle.Cors(handlers.InitialTransfer))

	server.HandleFunc(
		"/api/transfer",
		midle.Cors(midle.CheckSessionAndToken(handlers.LoadAllTransfer)))

	server.HandleFunc(
		"/unsecureGet",
		midle.Cors(handlers.LoadAllTransfer))

	server.HandleFunc(
		"GET /api/checkCookie",
		midle.Cors(handlers.CheckCookie))

	server.HandleFunc(
		"/acceptTransfer",
		midle.Cors(midle.CheckJwt(midle.IsAdmin(handlers.AdminAcceptTransfer))))

	server.HandleFunc(
		"/api/resetPassword", //zmienić potem tylko na post, na razie są problemy z CORS
		midle.Cors(handlers.ResetPass))

	//wysyłanie i potwoerdzanie przelewu
	server.HandleFunc(
		"POST /saveTransaction",
		midle.Cors(midle.CheckSessionAndToken(handlers.SaveTransaction)))

	server.HandleFunc(
		"/initTransaction", //zmienić potem tylko na post, na razie są problemy z CORS
		midle.Cors(midle.CheckSessionAndToken(handlers.InitialTransfer)))

	server.HandleFunc(
		"GET /saveTransaction", //zmienić potem tylko na post, na razie są problemy z CORS
		midle.Cors(midle.CheckToken(handlers.InitialTransferAccept)))

	server.HandleFunc(
		"/api/setSessionId",
		midle.Cors(handlers.CheckCookie))

	server.HandleFunc(
		"/essaTest",
		handlers.ThreadTest,
	)

	//http.ListenAndServe("127.0.0.1:8080", server)
	err = http.ListenAndServeTLS("127.0.0.1:8080",
		"/home/dalinarkholin/Uczelnia/wstepDoBezpieczenstwa/jebacKlucze/XD/example.com.crt",
		"/home/dalinarkholin/Uczelnia/wstepDoBezpieczenstwa/jebacKlucze/XD/example.com.key",
		server)
	if err != nil {
		log.Fatalf("ListenAndServeTLS error: %v", err)
	}
}

func xd() {
	cert, err := tls.LoadX509KeyPair("server-cert.pem", "server-key.pem")
	if err != nil {
		log.Fatalf("failed to load server certificate and key: %s", err)
	}

	// Load CA certificate
	caCert, err := os.ReadFile("ca-cert.pem")
	if err != nil {
		log.Fatalf("failed to read CA certificate: %s", err)
	}

	// Create a CA certificate pool and add the CA certificate to it
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Create a TLS configuration
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientCAs:    caCertPool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
	}

	// Create an HTTP server with the TLS configuration
	server := &http.Server{
		Addr:      ":8081",
		TLSConfig: tlsConfig,
	}

	log.Println("Starting server on https://localhost:8081")
	err = server.ListenAndServeTLS("", "")
	if err != nil {
		log.Fatalf("server failed to start: %s", err)
	}
}

//138
