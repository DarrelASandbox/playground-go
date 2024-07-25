package webserver

import (
	"embed"
	"html/template"
	"net/http"
)

//go:embed markup/*.gohtml
var templates embed.FS

type handler struct {
	templ *template.Template
}

func NewHandler() (http.Handler, error) {
	templ, err := template.ParseFS(templates, "markup/*.gohtml")
	if err != nil {
		return nil, err
	}

	handler := handler{templ: templ}
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.form)
	return mux, nil
}

func (h handler) form(w http.ResponseWriter, _ *http.Request) {
	_ = h.templ.ExecuteTemplate(w, "form.gohtml", nil)
}
