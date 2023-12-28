package main

import (
	"clothesPicker/server"
	"log"
	"math/rand"
	"net/http"
)

type InMemoryClothesStore struct {
	clothes server.Clothes
}

func (i *InMemoryClothesStore) RecordNewClothes(c server.Clothes) {}

func (i *InMemoryClothesStore) GetRandomClothing() string {
	return i.clothes[rand.Intn(len(i.clothes))]
}

func (i *InMemoryClothesStore) GetAllClothes() server.Clothes {
	return i.clothes
}

func main() {
	clothes := server.Clothes{
		"blue jeans",
		"blue sweater",
		"red hoodie",
	}
	store := InMemoryClothesStore{clothes}
	server := server.NewClothesServer(&store)

	log.Fatal(http.ListenAndServe("localhost:5000", server))
}
