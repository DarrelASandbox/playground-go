package blogposts

import (
	"io"
)

type Post struct {
	Title       string
	Description string
}

/*
In our case we only use it as an argument to io.ReadAll which needs an io.Reader.
So we should loosen the coupling in our function and ask for an io.Reader.
*/
func newPost(postFile io.Reader) (Post, error) {
	postData, err := io.ReadAll(postFile)
	if err != nil {
		return Post{}, err
	}

	post := Post{Title: string(postData)[7:]}
	return post, nil
}
