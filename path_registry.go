package main

import (
	"net/http"
)

func registerPaths(mux *http.ServeMux) {
	mux.HandleFunc("GET /v1/readiness", handlerReadiness)
	mux.HandleFunc("GET /v1/err", handlerErr)

}
