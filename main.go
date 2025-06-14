package main

import (
	"log"
	"net/http"
)

func songsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from songs!"))
}

func main() {
	http.HandleFunc("/songs", songsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
