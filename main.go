package main

import (
	"embed"
	"html/template"
	"net/http"
	"os"
)

//go:embed template/* static/*
var webFS embed.FS

type PageData struct {
	Result   template.HTML
	Text     string
	Banner   string
	Color    string
	SubMatch string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Load templates straight from the embedded memory file system
	tmpl, err := template.ParseFS(webFS, "template/index.html")
	if err != nil {
		http.Error(w, "Template Component Missing", http.StatusNotFound)
		return
	}

	tmpl.Execute(w, PageData{})
}

func asciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	userInput := r.FormValue("text")
	userBanner := r.FormValue("banner")
	chosenColor := r.FormValue("color")
	targetSubstring := r.FormValue("subMatch")

	ascii, err := printAscii(userBanner, userInput, chosenColor, targetSubstring)
	if err != nil {
		http.Error(w, "Internal Generation Fault", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFS(webFS, "template/index.html")
	if err != nil {
		http.Error(w, "Template Component Missing", http.StatusNotFound)
		return
	}

	data := PageData{
		Result:   template.HTML(ascii),
		Text:     userInput,
		Banner:   userBanner,
		Color:    chosenColor,
		SubMatch: targetSubstring,
	}

	tmpl.Execute(w, data)
}

func main() {
	// Feed your embedded static assets directory to the HTTP FileServer router
	http.Handle("/static/", http.FileServer(http.FS(webFS)))

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiHandler)

	// Vercel and dynamic cloud host systems inject the port dynamically via an environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.ListenAndServe(":"+port, nil)
}
