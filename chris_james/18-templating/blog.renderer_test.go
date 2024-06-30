package renderer_test

import (
	"bytes"
	"io"
	"testing"

	blog_renderer "github.com/DarrelASandbox/playground-go/chris_james/18-templating"
	approvals "github.com/approvals/go-approval-tests"
)

func TestReader(t *testing.T) {
	var (
		aPost = blog_renderer.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	postRenderer, err := blog_renderer.NewPostRenderer()

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
}

// BenchmarkRender-8          33313             33963 ns/op
// BenchmarkRender-8         489895              2414 ns/op
func BenchmarkRender(b *testing.B) {
	var (
		aPost = blog_renderer.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	postRenderer, err := blog_renderer.NewPostRenderer()

	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		postRenderer.Render(io.Discard, aPost)
	}
}
