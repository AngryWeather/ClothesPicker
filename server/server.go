package server

import (
	"encoding/json"
	"net/http"
)

type ClothesStore interface {
	GetRandomClothing() Clothes
}

type ClothesServer struct {
	Store ClothesStore
}

type Clothes struct {
	Name string
}

func (c *ClothesServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(c.Store.GetRandomClothing())
}
