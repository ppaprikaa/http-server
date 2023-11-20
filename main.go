package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	host := ""
	port := 5000
	addr := fmt.Sprintf("%s:%d", host, port)
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
