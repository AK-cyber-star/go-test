package bar

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlePostBAR(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(HandlePostBAR))
	defer server.Close()

	bar := &Bar{
		Name: "sam",
		Age:  20,
	}

	jsonData, err := json.Marshal(bar)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := http.Post(server.URL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected %d, but got %d instead.", http.StatusOK, resp.StatusCode)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var responseBar Bar
	if err := json.Unmarshal(b, &responseBar); err != nil {
		t.Fatal(err)
	}

	if *bar != responseBar {
		t.Errorf("expected %s, but got %s instead.", string(jsonData), string(b))
	}
}
