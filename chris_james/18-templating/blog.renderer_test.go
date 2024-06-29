package renderer_test

import (
	"bytes"
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
	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := blog_renderer.Render(&buf, aPost)

		if err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}
