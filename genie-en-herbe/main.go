package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	// "os/exec"
)

// clearCmd := exec.Command("clear")
// clearCmd.Stdout = os.Stdout
// clearCmd.Run()
// fmt.Println(message)

type Player struct {
	id string
	name string
	xp int
}

type Question struct {
	id int
	question string
	response string
	point int
}
 
var f = fmt.Println 
var s = fmt.Scan
var format = fmt.Sprintf

func main() {

	content, err := os.ReadFile("./questions.txt")

	if err != nil {
		f(err)
		os.Exit(0)
	}
	
	_questions := strings.Split(string(content), "\n")
	n := len(_questions)

	allQuestions := make([]Question, n)

	for i, v := range _questions {
		line := strings.Split(v, " : ")
		allQuestions[i].id = i
		allQuestions[i].question = line[0]
		allQuestions[i].response = line[1]
		point, _ := strconv.Atoi(line[2])
		allQuestions[i].point = point
	}

	var nb_joueurs int

	f("Combien de joueurs vous avez ?")
	s(&nb_joueurs)
	
	Players := make([]Player, nb_joueurs)

	lines := ""

	for i, _ := range Players {
		Players[i].id = strconv.Itoa(i) + "A"
		f("Donnez le nom du joueur ", i+1)
		s(&Players[i].name)
		lines += format("%v %v %v\n", Players[i].id, Players[i].name, Players[i].xp)
	}

	os.WriteFile("./players.txt", []byte(lines), 0666)
	
	for i:=0;i<len(allQuestions);i++ {
		badNameOrAnswer:
			f("Interrogez Un Joueur: ")
			name := ""
			s(&name)
			if isIn(name, Players) {
				f(allQuestions[i].question)
				f("Sa réponse, Vrai ou Faux\nV pour Vrai F pour Faux");var r rune;s(&r)
				if r != 'V' {
					currentPlayer := getPlayer(name, Players)
					f(currentPlayer)
				} else {
					f("Bad Answer")
					goto badNameOrAnswer
				}
			} else {
				f("Bad Name")
				goto badNameOrAnswer
			}
	}
}

func getId(name string, Players []Player) string {
	for i:=0;i<len(Players);i++ {
		if Players[i].name == name {
			return Players[i].id
		}
	}
	return ""
}