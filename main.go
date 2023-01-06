package main

import (
	"fmt"
	"log"
	"net/http"
)

// home page handler
func homePage(w http.ResponseWriter, r *http.Request) {
	html := `<strong>Hello, world</strong>`
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, html)
}

func main() {
	http.HandleFunc("/", homePage)

	log.Println("starting web server...")
	http.ListenAndServe(":8080", nil)
}
