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

func (s *StubClothesStore) GetAllClothes() []Clothes {
	return s.clothes
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

		var got string
		err := json.NewDecoder(response.Body).Decode(&got)

		if err != nil {
			t.Fatalf("Unable to parse response from server %q into slice of %v", response.Body, err)
		}

		want := "blue sweater"

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %+v, want %+v", got, want)
		}

		if response.Result().Header.Get("content-type") != "application/json" {
			t.Errorf("response did not have content-type of application/json, got %v", response.Result().Header)
		}
	})
}

// func TestGetAllClothes(t *testing.T) {
// 	clothes := []Clothes{
// 		{Name: "blue jeans"},
// 		{Name: "blue sweater"},
// 	}

// 	store := StubClothesStore{
// 		clothes,
// 	}

// 	server := &ClothesServer{&store}

// 	t.Run("returns all clothes", func(t *testing.T) {
// 		request, _ := http.NewRequest(http.MethodGet, "/clothes", nil)
// 		response := httptest.NewRecorder()

// 		server.ServeHTTP(response, request)

// 		got := response.Result().StatusCode
// 		want := 200

// 		var clothes []Clothes
// 		err := json.NewDecoder(response.Body).Decode(&clothes)

// 		if err != nil {
// 			t.Fatalf("Unable to parse response from server %q into slice of %v", response.Body, err)
// 		}

// 		if got != want {
// 			t.Errorf("got code %d, want %d", got, want)
// 		}

// 		wantedClothes := []Clothes{
// 			{Name: "blue jeans"},
// 			{Name: "blue sweater"},
// 		}

// 		if !reflect.DeepEqual(wantedClothes, clothes) {
// 			t.Errorf("wanted %+v, got %+v", wantedClothes, clothes)
// 		}
// 	})
// }
