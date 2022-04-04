package main

import (
	"Api-Queries/server"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	server.New(mux)
	http.ListenAndServe(":8080", mux)
}
