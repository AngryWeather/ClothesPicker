package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRandomClothing(t *testing.T) {
	t.Run("returns random clothing", func(t *testing.T) {
		clothes := []Clothes{
			{Name: "blue sweater"},
			{Name: "red hoodie"},
			{Name: "blue jeans"},
		}
		request, _ := http.NewRequest(http.MethodGet, "/random/clothing", nil)
		response := httptest.NewRecorder()

		var got []Clothes
		ClothingServer(response, request, clothes)

		err := json.NewDecoder(response.Body).Decode(&got)

		if err != nil {
			t.Fatalf("Unable to parse response from server %q into slice of %v", response.Body, err)
		}

		want := 1

		if len(got) != want {
			t.Errorf("got %d, want %d", len(got), want)
		}
	})

}
