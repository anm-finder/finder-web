package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	urlPath := r.URL.Path[1:]

	fmt.Fprintf(w, "Hi there, I love %s!", urlPath)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
