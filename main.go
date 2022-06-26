package main

import (
	"log"
	"net/http"
)

func main() {
	m := http.NewServeMux()

	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})

	s := http.Server{
		Addr:    ":8080",
		Handler: m,
	}

	log.Fatal(s.ListenAndServe())
}
