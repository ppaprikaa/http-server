package main

import (
	"log"
	"net/http"
)

func main() {
	addr := ":5000"
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	srv := http.Server{
		Addr:    addr,
		Handler: mux,
	}

	log.Printf("server is running on %s", addr)
	log.Fatal(srv.ListenAndServe())
}
