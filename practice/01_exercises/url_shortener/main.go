package main

import (
	"log"
	"net/http"
)

func main() {
	srv := &Server{store: NewMemoryStore(), baseURL: "http://localhost:8080"}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", srv.health)
	mux.HandleFunc("POST /links", srv.createLink)
	mux.HandleFunc("GET /links/{code}", srv.getLink)
	mux.HandleFunc("DELETE /links/{code}", srv.deleteLink)
	mux.HandleFunc("GET /{code}", srv.redirect)

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}