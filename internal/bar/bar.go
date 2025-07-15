package bar

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Bar struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func JSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func HandlePostBAR(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Println("invalid request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var bar Bar
	err := json.NewDecoder(r.Body).Decode(&bar)
	if err != nil {
		fmt.Println("Invalid data format")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	JSON(w, http.StatusOK, &bar)
}
