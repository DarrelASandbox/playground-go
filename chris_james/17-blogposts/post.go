package blogposts

import (
	"bufio"
	"io"
)

type Post struct {
	Title       string
	Description string
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
)

/*
`bufio.Scanner` scans through data, line by line
Separating the what from the how of reading lines to make the code a little more declarative to the reader.
*/
func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readLine := func() string {
		scanner.Scan()
		return scanner.Text()
	}

	title := readLine()[len(titleSeparator):]
	description := readLine()[len(descriptionSeparator):]

	return Post{Title: title, Description: description}, nil
}
