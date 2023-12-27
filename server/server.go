package server

import (
	"encoding/json"
	"net/http"
)

type ClothesStore interface {
	GetRandomClothing() string
	GetAllClothes() Clothes
}

type ClothesServer struct {
	Store ClothesStore
}

type Clothes []string

func (c *ClothesServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router := http.NewServeMux()

	w.Header().Set("content-type", "application/json")
	router.Handle("/random/clothes", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(c.Store.GetRandomClothing())
	}))

	router.Handle("/clothes", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(c.Store.GetAllClothes())
	}))

	router.ServeHTTP(w, r)
}
