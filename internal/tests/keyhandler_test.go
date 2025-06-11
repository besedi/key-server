package tests

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/besedi/key-server/internal/metrics"
	"github.com/besedi/key-server/internal/srv"
)

const maxSize = 100

func init() {
	metrics.Init(maxSize)
}

func TestKeyHandler_ValidLength(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/key/42", nil)
	req.SetPathValue("len", "42")

	rr := httptest.NewRecorder()
	handler := srv.KeyHandler(maxSize)
	handler.ServeHTTP(rr, req)

	if rr.Result().StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %d", rr.Result().StatusCode)
	}
}

func TestKeyHandler_ValidKeySize(t *testing.T) {
	for i := 1; i <= maxSize; i++ {
		l := strconv.Itoa(i)
		req := httptest.NewRequest(http.MethodGet, "/key/"+l, nil)
		req.SetPathValue("len", l)

		rr := httptest.NewRecorder()
		handler := srv.KeyHandler(maxSize)
		handler.ServeHTTP(rr, req)

		body, _ := io.ReadAll(rr.Body)
		bodyLen := len(body) / 8
		if i != bodyLen {
			t.Errorf("Expected len %d, got %d", i, bodyLen)
		}
	}

}

func TestKeyHandler_InvalidLength(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/key/999", nil) // exceeds maxSize
	req.SetPathValue("len", "999")
	rr := httptest.NewRecorder()

	handler := srv.KeyHandler(maxSize)
	handler.ServeHTTP(rr, req)

	res := rr.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status 400 Bad Request, got %d", res.StatusCode)
	}
}
