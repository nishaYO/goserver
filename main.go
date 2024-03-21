package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	m := http.NewServeMux()
	corsMux := middlewareCors(m)

	m.HandleFunc("/", servePage)
	server := http.Server{
		Handler:      corsMux,
		Addr:         ":8000",
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}
	fmt.Println("Server is listening at: localhost:8080")
	err := server.ListenAndServe()
	log.Fatal(err)
}

func servePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.WriteHeader(200)
	w.Write([]byte("hello this is nice text yoooo"))
	fmt.Println("Now serving the content over get req")
}

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			fmt.Println("preflight req made by the client")
			return
		}
		next.ServeHTTP(w, r)
	})
}