package main

import (
	"net/http"

	"github.com/syumai/workers"
	"github.com/zztkm/zztkm-bbs/app"
)

func main() {
	http.HandleFunc("/", app.HelloHandler)
	workers.Serve(nil) // use http.DefaultServeMux
}
