package main

import (
	"clothesPicker/server"
	"log"
	"net/http"
)

type InMemoryClothesStore struct{}

func (i *InMemoryClothesStore) GetRandomClothing() server.Clothes {
	return server.Clothes{}
}

func main() {
	server := &server.ClothesServer{&InMemoryClothesStore{}}

	log.Fatal(http.ListenAndServe("localhost:5000", server))
}
