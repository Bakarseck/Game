package handlers

import (
	"net/http"
	"html/template"
	"os"
)

type Data struct {
	Player1Name string
	Player2Name string
	score1 int
	score2 int
}

func MainHandler(w http.ResponseWriter, r *http.Request) {

	template, err := template.ParseFiles("./templates/index.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		os.Exit(0)
	}

	template.Execute(w, nil)
}

func PlayHandler(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./templates/play.html")

	e := r.ParseForm()

	if e != nil {
		os.Exit(0)
	}

	name1 := r.Form.Get("player1")
	name2 := r.Form.Get("player2")

	v := Data{
		Player1Name: name1,
		Player2Name: name2,
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		os.Exit(0)
	}

	template.Execute(w, v)
}