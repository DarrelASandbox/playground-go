package blogposts

import (
	"io/fs"
)

type Post struct {
}

/*
Even though our tests are passing, we can't use our new package outside of this context,
because it is coupled to a concrete implementation `fstest.MapFS`.
But, it doesn't have to be.
Change the argument to our `NewPostsFromFS` function to accept the interface from the standard library.

`fs.FS`: An interface for representing read-only file systems in a generic way, allowing for different implementations.
`fstest.MapFS`: A specific implementation of fs.FS designed for testing, providing an in-memory file system.
*/
func NewPostsFromFS(filesystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(filesystem, ".")
	if err != nil {
		return nil, err
	}

	var posts []Post
	for range dir {
		posts = append(posts, Post{})
	}
	return posts, nil
}
