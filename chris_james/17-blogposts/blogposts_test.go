/*
Notice that the package of our test is `blogposts_test`.
Remember, when TDD is practiced well we take a consumer-driven approach:
we don't want to test internal details because consumers don't care about them.
By appending `_test` to our intended package name, we only access exported members from our package
- just like a real user of our package.
*/
package blogposts_test

import (
	"testing"
	"testing/fstest"

	blogposts "github.com/DarrelASandbox/playground-go/chris_james/17-blogposts"
)

// A MapFS is a simple in-memory file system for use in tests, represented as a map from path names (arguments to Open) to information about the files or directories they represent.
func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("hi")},
		"hello-world2.md": {Data: []byte("hola")},
	}

	posts := blogposts.NewPostsFromFS(fs)

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}
}
