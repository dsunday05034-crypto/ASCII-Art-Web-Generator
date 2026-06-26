package main

import (
	"log"
	"net/http"
	"os"

	handler "ascii-art-web-generator/api"
)

func main() {
	// 1. Explicitly serve static assets locally when a path starts with "/static/"
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// 2. Map all root/application traffic to your Vercel logic function handler
	http.HandleFunc("/", handler.Handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Local dev server running smoothly on http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
