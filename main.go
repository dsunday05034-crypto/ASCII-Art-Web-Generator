package main

import (
	handler "ascii-art-web-generator/api"
	"log"
	"net/http"
	"os"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Map all root/application traffic to Vercel logic function handler
	http.HandleFunc("/", handler.Handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Local dev server running smoothly on http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
