package blogposts

import (
	"io"
	"io/fs"
)

type Post struct {
	Title string
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
	for _, f := range dir {
		post, err := getPost(filesystem, f)
		if err != nil {
			return nil, err // @TODO: needs clarification, should we totally fail if one file fails? or just ignore?
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func getPost(fileSystem fs.FS, f fs.DirEntry) (Post, error) {
	postFile, err := fileSystem.Open(f.Name())
	if err != nil {
		return Post{}, err
	}

	defer postFile.Close()

	postData, err := io.ReadAll(postFile)
	if err != nil {
		return Post{}, err
	}

	post := Post{Title: string(postData)[7:]}
	return post, nil
}
