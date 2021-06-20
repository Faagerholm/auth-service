package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", PingHandler)
	r.HandleFunc("/login", LoginHandler)
	r.HandleFunc("/logout", LogoutHandler)
	r.HandleFunc("/check-token", CheckTokenHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:6000",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	log.Println("Listening on incoming requests")
	log.Fatal(srv.ListenAndServe())
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Pong")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "User Logged in")
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Logged out successfully")
}

func CheckTokenHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusForbidden)
	fmt.Fprint(w, "User not authenticated")
}
