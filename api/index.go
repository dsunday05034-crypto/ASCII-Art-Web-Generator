package api // Matches your folder workspace architecture cleanly

import (
	"embed"
	"net/http"
	"text/template"
)

//go:embed template/index.html
var templateFS embed.FS

type PageData struct {
	Text     string
	Banner   string
	Color    string
	SubMatch string
	Result   string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("index.html").ParseFS(templateFS, "template/index.html")
	if err != nil {
		http.Error(w, "Template compilation error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodGet {
		tmpl.Execute(w, PageData{})
		return
	}

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Form parsing error", http.StatusBadRequest)
			return
		}

		text := r.FormValue("text")
		banner := r.FormValue("banner")
		color := r.FormValue("color")
		subMatch := r.FormValue("subMatch")

		// Calls the engine logic sitting within the same package seamlessly
		asciiOutput, err := PrintAscii(banner, text, color, subMatch)
		if err != nil {
			http.Error(w, "ASCII processing failure: "+err.Error(), http.StatusInternalServerError)
			return
		}

		data := PageData{
			Text:     text,
			Banner:   banner,
			Color:    color,
			SubMatch: subMatch,
			Result:   asciiOutput,
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		tmpl.Execute(w, data)
		return
	}

	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}
