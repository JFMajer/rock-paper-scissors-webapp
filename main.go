package main

import (
	"fmt"
	"html/template"
	"log"
	"myapp/rps"
	"net/http"
)

const portNumber = ":8080"

// home page handler
func homePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

func playRound(w http.ResponseWriter, r *http.Request) {
	winner, computerChoice, roundResult := rps.PlayRound(1)
	log.Println(winner, computerChoice, roundResult)
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
	}
}

func main() {
	http.HandleFunc("/play", playRound)
	http.HandleFunc("/", homePage)

	fmt.Printf("Starting application on port %s\n", portNumber)
	http.ListenAndServe(portNumber, nil)
}
