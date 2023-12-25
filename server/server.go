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
	var randomC = []Clothes{c[rand.Intn(3)]}
	json.NewEncoder(w).Encode(randomC)
}
