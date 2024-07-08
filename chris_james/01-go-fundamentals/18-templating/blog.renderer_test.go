package renderer_test

import (
	"bytes"
	"io"
	"testing"

	blogRenderer "github.com/DarrelASandbox/playground-go/chris_james/18-templating"
	approvals "github.com/approvals/go-approval-tests"
)

func TestReader(t *testing.T) {
	var (
		aPost = blogRenderer.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	postRenderer, err := blogRenderer.NewPostRenderer()

	if err != nil {
		t.Fatal(err)
	}

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := postRenderer.Render(&buf, aPost)

		if err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})

	/*
	   We're using the `Post`'s title field as a part of the path of the URL,
	   but we don't really want spaces in the URL so we're replacing them with hyphens.
	   We've added a `RenderIndex` method to our `PostRenderer` that again takes an `io.Writer` and a slice of `Post`.
	*/
	t.Run("it renders an index of posts", func(t *testing.T) {
		buf := bytes.Buffer{}
		posts := []blogRenderer.Post{{Title: "Hello World"}, {Title: "Hello World 2"}}

		if err := postRenderer.RenderIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}

// BenchmarkRender-8          33313             33963 ns/op
// BenchmarkRender-8         489895              2414 ns/op
func BenchmarkRender(b *testing.B) {
	var (
		aPost = blogRenderer.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	postRenderer, err := blogRenderer.NewPostRenderer()

	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		postRenderer.Render(io.Discard, aPost)
	}
}
