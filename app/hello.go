package app

import (
	"fmt"
	"net/http"
	"os"

	"github.com/mailru/easyjson"
)

//easyjson:json
type HelloResponse struct {
	Message string `json:"message"`
}

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	msg := "Hello from syumai/workers!"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res := HelloResponse{Message: msg}
	if _, err := easyjson.MarshalToWriter(&res, w); err != nil {
		fmt.Fprintf(os.Stderr, "failed to encode response: %w\n", err)
	}
}
