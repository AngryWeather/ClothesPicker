package main

import (
	"bytes"
	"clothesPicker/server"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestRecordingClothesAndRetrieving(t *testing.T) {
	store := InMemoryClothesStore{}
	sv := server.NewClothesServer(&store)

	sv.ServeHTTP(httptest.NewRecorder(), newClothesRequest("blue jeans"))
	sv.ServeHTTP(httptest.NewRecorder(), newClothesRequest("red hoodie"))
	sv.ServeHTTP(httptest.NewRecorder(), newClothesRequest("black jeans"))

	response := httptest.NewRecorder()
	sv.ServeHTTP(response, newGetAllClothesRequest())

	if response.Code != http.StatusOK {
		t.Errorf("got status %d, want %d", response.Code, http.StatusOK)
	}

	var clothes server.Clothes

	err := json.NewDecoder(response.Body).Decode(&clothes)

	if err != nil {
		t.Fatalf("Unable to parse response from %v", clothes)
	}

	want := server.Clothes{"blue jeans", "red hoodie", "black jeans"}
	if !reflect.DeepEqual(clothes, want) {
		t.Errorf("got %v, want %v", response.Body, want)
	}
}

func newGetAllClothesRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/clothes/", nil)
	return req
}

func newClothesRequest(name string) *http.Request {
	requestBody, _ := json.Marshal(name)
	request, _ := http.NewRequest(http.MethodPost, "/clothes/", bytes.NewBuffer(requestBody))

	return request
}
