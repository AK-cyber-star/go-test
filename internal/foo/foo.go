package foo

import (
	"fmt"
	"net/http"
)

func HandleGETFOO(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		fmt.Println("invalid request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("FOO"))
}
