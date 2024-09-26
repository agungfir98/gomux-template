package handler

import (
	"encoding/json"
	"net/http"
)

type Result struct {
	Message string `json:"message"`
}

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	example := Result{
		Message: "Halo sekai",
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(example)
	return
}
