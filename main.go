package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func songsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World"))
}

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/songs", songsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
