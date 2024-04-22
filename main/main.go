package main

import (
	"GoGames/backend/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/api/checkword", handlers.WordCheck)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
