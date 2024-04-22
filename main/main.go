package main

import (
	"GoGames/backend/games/sudoku"
	"fmt"
	"log"
	"net/http"
)

func home_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Homepage")
}

func handle_requests() {
	http.HandleFunc("/", home_page)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	sudoku.DrawCells()
	fmt.Println("Environment Test")
	handle_requests()
}
