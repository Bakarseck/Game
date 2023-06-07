package main

import (
	"net/http"
	"log"

	"game/scrabble-note/internals/handlers"
)



func main() {

	http.HandleFunc("/", handlers.MainHandler)
	http.HandleFunc("/play", handlers.PlayHandler)

	// fileServers := http.ServeFiles(http.Dir("./static/"))
	// http.Handle("/static/", fileServers)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	log.Println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}