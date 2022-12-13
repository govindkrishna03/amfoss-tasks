package main

import (
	"log"
	"net/http"
)

func main() {

	err := http.ListenAndServe(":9000", http.FileServer(http.Dir("../../assets")))
	log.Fatalln(err)
}
