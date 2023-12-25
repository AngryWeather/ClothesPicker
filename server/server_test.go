package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRandomClothing(t *testing.T) {
	t.Run("returns clothing types", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/categories", nil)
		response := httptest.NewRecorder()

		ClothingServer(response, request)

		got := response.Body.String()
		want := "blue sweater"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
