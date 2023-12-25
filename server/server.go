package server

import (
	"fmt"
	"net/http"
)

func ClothingServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "blue sweater")
}
