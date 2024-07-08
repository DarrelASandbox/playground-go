package blogposts

import (
	"io/fs"
)

/*
Even though our tests are passing, we can't use our new package outside of this context,
because it is coupled to a concrete implementation `fstest.MapFS`.
But, it doesn't have to be.
Change the argument to our `NewPostsFromFS` function to accept the interface from the standard library.

`fs.FS`: An interface for representing read-only file systems in a generic way, allowing for different implementations.
`fstest.MapFS`: A specific implementation of fs.FS designed for testing, providing an in-memory file system.
*/
func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}

	var posts []Post
	for _, f := range dir {
		post, err := getPost(fileSystem, f.Name())
		if err != nil {
			return nil, err // @TODO: needs clarification, should we totally fail if one file fails? or just ignore?
		}

		posts = append(posts, post)
	}

	return posts, nil
}

/*
When you refactor out new functions or methods, take care and think about the arguments.
You're designing here, and are free to think deeply about what is appropriate because you have passing tests.
Think about coupling and cohesion. In this case you should ask yourself:

Does newPost have to be coupled to an `fs.File`?
Do we use all the methods and data from this type?
What do we really need?
*/
func getPost(fileSystem fs.FS, fileName string) (Post, error) {
	postFile, err := fileSystem.Open(fileName)
	if err != nil {
		return Post{}, err
	}

	defer postFile.Close()
	return newPost(postFile)
}
