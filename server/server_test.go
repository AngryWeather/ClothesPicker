package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type StubClothesStore struct {
	clothes []Clothes
}

func (s *StubClothesStore) GetRandomClothing() Clothes {
	return s.clothes[1]
}

func TestRandomClothing(t *testing.T) {
	clothes := []Clothes{
		{Name: "blue jeans"},
		{Name: "blue sweater"},
	}

	store := StubClothesStore{
		clothes,
	}

	server := &ClothesServer{&store}

	t.Run("returns random clothing as length of 1", func(t *testing.T) {

		request, _ := http.NewRequest(http.MethodGet, "/random/clothes", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		var got Clothes
		err := json.NewDecoder(response.Body).Decode(&got)

		if err != nil {
			t.Fatalf("Unable to parse response from server %q into slice of %v", response.Body, err)
		}

		want := Clothes{Name: "blue sweater"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %+v, want %+v", got, want)
		}
	})

}
