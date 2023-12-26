package main

import (
	"clothesPicker/server"
	"log"
	"math/rand"
	"net/http"
)

type InMemoryClothesStore struct {
	clothes []server.Clothes
}

func (i *InMemoryClothesStore) GetRandomClothing() server.Clothes {
	return i.clothes[rand.Intn(len(i.clothes))]
}

func (i *InMemoryClothesStore) GetAllClothes() []server.Clothes {
	return i.clothes
}

func main() {
	clothes := []server.Clothes{
		{Name: "blue jeans"},
		{Name: "blue sweater"},
		{Name: "red hoodie"},
	}
	store := InMemoryClothesStore{clothes}
	server := &server.ClothesServer{Store: &store}

	log.Fatal(http.ListenAndServe("localhost:5000", server))
}
