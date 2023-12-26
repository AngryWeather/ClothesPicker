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
	w.Header().Set("content-type", "application/json")
	if r.URL.Path == "/random/clothes" {
		json.NewEncoder(w).Encode(c.Store.GetRandomClothing())
	} else if r.URL.Path == "/clothes" {
		json.NewEncoder(w).Encode([]Clothes{
			{Name: "blue jeans"},
			{Name: "blue sweater"},
		})
	}
}
