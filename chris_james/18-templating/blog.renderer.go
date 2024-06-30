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

type PostRenderer struct {
	templ *template.Template
}

//go:embed "templates/*"
var postTemplates embed.FS

func Render(w io.Writer, p Post) error {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return err
	}

	if err := templ.ExecuteTemplate(w, "blog.gohtml", p); err != nil {
		return err
	}

	return err
}

// To stop us having to re-parse the templates over and over,
// we'll create a type that'll hold the parsed template, and that'll have a method to do the rendering
func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	return &PostRenderer{templ: templ}, nil
}

func (r *PostRenderer) Render(w io.Writer, p Post) error {
	if err := r.templ.ExecuteTemplate(w, "blog.gohtml", p); err != nil {
		return err
	}

	return nil
}

/*
FAIL: TestReader/it_renders_an_index_of_posts (0.00s)

	        blog.renderer_test.go:56: got
					"<ol><li><a href=\"/post/Hello%20World\">Hello World</a></li><li><a href=\"/post/Hello%20World%202\">Hello World 2</a></li></ol>"
					want
					"<ol><li><a href=\"/post/hello-world\">Hello World</a></li><li><a href=\"/post/hello-world-2\">Hello World 2</a></li></ol>"
*/

/*
You can see that the templating code is escaping the spaces in the `href` attributes.
We need a way to do a string replace of spaces with hyphens.
We can't just loop through the `[]Post` and replace them in-memory because
we still want the spaces displayed to the user in the anchors.
*/
func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	indexTemplate := `<ol>{{range .}}<li><a href="/post/{{.Title}}">{{.Title}}</a></li>{{end}}</ol>`

	templ, err := template.New("index").Parse(indexTemplate)
	if err != nil {
		return err
	}

	if err := templ.Execute(w, posts); err != nil {
		return err
	}

	return nil
}
