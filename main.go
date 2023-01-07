package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"myapp/rps"
	"net/http"
	"strconv"
)

const portNumber = ":8080"

// home page handler
func homePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

func playRound(w http.ResponseWriter, r *http.Request) {
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("c"))
	result := rps.PlayRound(playerChoice)
	out, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)

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
