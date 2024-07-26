package webserver

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/specs-greet/domain/interactions"
)

const (
	greetPath = "/greet"
	cursePath = "/curse"
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
	mux.HandleFunc(greetPath, handler.replyWith(interactions.Greet))
	mux.HandleFunc(cursePath, handler.replyWith(interactions.Curse))
	return mux, nil
}

func (h handler) form(w http.ResponseWriter, _ *http.Request) {
	_ = h.templ.ExecuteTemplate(w, "form.gohtml", nil)
}

func (h handler) replyWith(interact func(name string) string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := h.templ.ExecuteTemplate(w, "reply.gohtml", interact(r.Form.Get("name"))); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
