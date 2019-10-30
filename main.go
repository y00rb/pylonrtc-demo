package main

import (
	"log"
	"net/http"
)

func serveHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/", serveHome)

	err := http.ListenAndServe("localhost:8282", nil)
	if err != nil {
		log.Fatal("Error on ListenAndServe", err)
	}
}
