package main

import (
	"log"
	"net/http"
)

func main() {
	server := &PlayerServer{NewInMemoryPlayerStore(), nil}

	log.Fatal(http.ListenAndServe(":8000", server))
}
