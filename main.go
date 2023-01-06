package main

import (
	"fmt"
	"log"
	"net/http"
)

// home page handler
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world")
}

func main() {
	http.HandleFunc("/", homePage)

	log.Println("starting web server...")
	http.ListenAndServe(":8080", nil)
}
