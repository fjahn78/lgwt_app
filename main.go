package main

import (
	"log"
	"net/http"
)

func main() {
	handler := &PlayerServer{}
	log.Fatal((http.ListenAndServe(":5000", handler)))
}
