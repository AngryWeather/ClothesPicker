package server

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type ClothesStore interface {
	GetRandomClothing() string
	GetAllClothes() Clothes
	RecordNewClothes(s string)
	GetClothesById(i int) string
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
	router.Handle("/clothes/", http.HandlerFunc(c.clothesHandler))

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
	id_prefix := strings.TrimPrefix(r.URL.Path, "/clothes/")

	if len(id_prefix) == 0 {
		switch r.Method {
		case http.MethodPost:
			w.WriteHeader(http.StatusAccepted)
			clothes := c.decodeClothesJson(r)
			c.Store.RecordNewClothes(clothes)
		case http.MethodGet:
			c.showClothes(w)
		}
	} else {
		id, _ := strconv.Atoi(id_prefix)
		c.getClothesById(w, id)
	}
}

func (c *ClothesServer) getClothesById(w http.ResponseWriter, id int) {
	json.NewEncoder(w).Encode(c.Store.GetClothesById(id))
}

func (c *ClothesServer) decodeClothesJson(r *http.Request) string {
	var clothes string
	err := json.NewDecoder(r.Body).Decode(&clothes)
	if err != nil {
		panic("could not parse json")
	}
	return clothes
}

func (c *ClothesServer) showClothes(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(c.Store.GetAllClothes())
}

func setJsonHeader(w http.ResponseWriter) {
	w.Header().Set("content-type", "application/json")
}
