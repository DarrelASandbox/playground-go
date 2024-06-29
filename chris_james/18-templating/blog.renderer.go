package renderer

import (
	"embed"
	"html/template"
	"io"
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

//go:embed "templates/*"
var postTemplates embed.FS

func Render(w io.Writer, p Post) error {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return err
	}

	if err := templ.Execute(w, p); err != nil {
		return err
	}

	return err
}
