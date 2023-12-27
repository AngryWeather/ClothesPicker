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
	Store  ClothesStore
	router *http.ServeMux
}

func NewClothesServer(store ClothesStore) *ClothesServer {
	c := &ClothesServer{
		store,
		http.NewServeMux(),
	}

	c.router.Handle("/random/clothes", http.HandlerFunc(c.randomClothesHandler))
	c.router.Handle("/clothes", http.HandlerFunc(c.clothesHandler))

	return c
}

type Clothes []string

func (c *ClothesServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	c.router.ServeHTTP(w, r)
}

func (c *ClothesServer) randomClothesHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(c.Store.GetRandomClothing())
}

func (c *ClothesServer) clothesHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(c.Store.GetAllClothes())
}
