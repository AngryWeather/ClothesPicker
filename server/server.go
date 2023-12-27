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
	http.Handler
}

func NewClothesServer(store ClothesStore) *ClothesServer {
	c := new(ClothesServer)
	c.Store = store

	router := http.NewServeMux()
	router.Handle("/random/clothes", http.HandlerFunc(c.randomClothesHandler))
	router.Handle("/clothes", http.HandlerFunc(c.clothesHandler))

	c.Handler = router

	return c
}

type Clothes []string

func (c *ClothesServer) randomClothesHandler(w http.ResponseWriter, r *http.Request) {
	setJsonHeader(w)
	json.NewEncoder(w).Encode(c.Store.GetRandomClothing())
}

func (c *ClothesServer) clothesHandler(w http.ResponseWriter, r *http.Request) {
	setJsonHeader(w)
	json.NewEncoder(w).Encode(c.Store.GetAllClothes())
}

func setJsonHeader(w http.ResponseWriter) {
	w.Header().Set("content-type", "application/json")
}
