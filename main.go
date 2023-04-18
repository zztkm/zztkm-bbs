package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/syumai/workers"
)

type HelloResponse struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		msg := "Hello!"
		w.Header().Set("Content-Type", "application/json")
		res := HelloResponse{Message: msg}
		if err := json.NewEncoder(w).Encode(res); err != nil {
			fmt.Fprintf(os.Stderr, "failed to encode response: %w\n", err)
		}
	})
	workers.Serve(nil) // use http.DefaultServeMux
}
