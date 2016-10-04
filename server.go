package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	var home = os.Getenv("HOME")
	if home == "" {
		home = "./"
	}
	var port = os.Getenv("PORT")
	if port == "" {
		port = "4747"
	}

	fs := http.FileServer(http.Dir(home + "/public/"))
	http.Handle("/", fs)

	log.Println("Listening...")
	http.ListenAndServe(":"+port, nil)
}
