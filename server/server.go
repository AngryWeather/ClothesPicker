package server

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

type Clothes struct {
	Name string
}

func ClothingServer(w http.ResponseWriter, r *http.Request, c []Clothes) {
	json.NewEncoder(w).Encode(GetRandomClothing(c))
}

func GetRandomClothing(c []Clothes) []Clothes {
	return []Clothes{c[rand.Intn(len(c))]}
}
