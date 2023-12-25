package main

import (
	"clothesPicker/server"
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(server.ClothingServer)
	log.Fatal(http.ListenAndServe("localhost:5000", handler))
}
