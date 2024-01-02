package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type StubClothesStore struct {
	clothes         Clothes
	newClothesCalls []string
}

func (s *StubClothesStore) GetClothesById(i int) string {
	return s.clothes[i-1]
}

func (s *StubClothesStore) GetRandomClothing() string {
	return s.clothes[1]
}

func (s *StubClothesStore) RecordNewClothes(str string) {
	s.newClothesCalls = append(s.newClothesCalls, str)
	print(s.newClothesCalls)
}

func (s *StubClothesStore) GetAllClothes() Clothes {
	return s.clothes
}

func TestRandomClothing(t *testing.T) {
	clothes := Clothes{
		"blue jeans",
		"blue sweater",
	}

	store := StubClothesStore{
		clothes,
		nil,
	}

	server := NewClothesServer(&store)

	t.Run("returns random clothing as string", func(t *testing.T) {

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

func TestGetAllClothes(t *testing.T) {
	clothes := Clothes{
		"blue jeans",
		"blue sweater",
	}

	store := StubClothesStore{
		clothes,
		nil,
	}

	server := NewClothesServer(&store)

	t.Run("returns all clothes", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/clothes/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Result().StatusCode
		want := 200

		var clothes Clothes
		err := json.NewDecoder(response.Body).Decode(&clothes)

		if err != nil {
			t.Fatalf("Unable to parse response from server %q into slice of %v", response.Body, err)
		}

		assertStatus(t, got, want)

		wantedClothes := Clothes{
			"blue jeans",
			"blue sweater",
		}

		if !reflect.DeepEqual(wantedClothes, clothes) {
			t.Errorf("wanted %+v, got %+v", wantedClothes, clothes)
		}
	})
}

func TestPostClothes(t *testing.T) {
	clothes := Clothes{
		"blue jeans",
		"blue sweater",
	}

	store := StubClothesStore{
		clothes,
		nil,
	}

	server := NewClothesServer(&store)

	t.Run("creates new clothing", func(t *testing.T) {
		requestBody, _ := json.Marshal("red trousers")
		request, _ := http.NewRequest(http.MethodPost, "/clothes/", bytes.NewBuffer(requestBody))
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Result().StatusCode
		want := 202

		assertStatus(t, got, want)

		if len(store.newClothesCalls) != 1 {
			t.Errorf("got %d calls to RecordNewClothes, want %d", len(store.newClothesCalls), 1)
		}

		if !reflect.DeepEqual(store.newClothesCalls[0], "red trousers") {
			t.Errorf("got %v, want %v", store.newClothesCalls[0], "red trousers")
		}
	})
}

func TestGetClothingWithId(t *testing.T) {
	clothes := Clothes{
		"blue jeans",
		"blue sweater",
	}

	store := StubClothesStore{
		clothes,
		nil,
	}

	server := NewClothesServer(&store)

	t.Run("returns clothing with id 1", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/clothes/1", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		var got string
		err := json.NewDecoder(response.Body).Decode(&got)

		if err != nil {
			t.Fatalf("Unable to parse response from server %q into slice of %v", response.Body, err)
		}

		want := "blue jeans"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns clothing with id 2", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/clothes/2", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		var got string
		err := json.NewDecoder(response.Body).Decode(&got)

		if err != nil {
			t.Fatalf("Unable to parse response from server %q into slice of %v", response.Body, err)
		}

		want := "blue sweater"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns 404 on missing id", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/clothes/3", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Result().StatusCode
		want := 404

		assertStatus(t, got, want)

	})
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got code %d, want %d", got, want)
	}
}
