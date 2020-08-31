package main

import (
	"log"
	"net/http"

	"./addidea"
	"./index"
	"./see"
)

func main() {
	http.HandleFunc("/", index.Index)
	http.HandleFunc("/add", addidea.AddIdea)
	http.HandleFunc("/see", see.See)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
