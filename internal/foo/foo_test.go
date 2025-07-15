package foo

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleFOO(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(HandleGETFOO))
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected %d, but got %d.", http.StatusOK, resp.StatusCode)
	}

	expected := "FOO"
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != expected {
		t.Errorf("expected %s, but got %s.", expected, string(b))
	}
}
