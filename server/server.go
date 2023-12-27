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

	router.Handle("/random/clothes", http.HandlerFunc(c.randomClothesHandler))
	router.Handle("/clothes", http.HandlerFunc(c.clothesHandler))

	router.ServeHTTP(w, r)
}

func (c *ClothesServer) randomClothesHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(c.Store.GetRandomClothing())
}

func (c *ClothesServer) clothesHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(c.Store.GetAllClothes())
}
