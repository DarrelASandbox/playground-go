package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagSeparator         = "Tags: "
)

/*
`bufio.Scanner` scans through data, line by line

Scan the next line to ignore the --- separator.
Keep scanning until there's nothing left to scan.

`scanner.Scan()` returns a `bool` which indicates whether there's more data to scan,
so we can use that with a `for` loop to keep reading through the data until the end.

After every `Scan()` we write the data into the buffer using `fmt.Fprintln`.
We use the version that adds a newline because the scanner removes the newlines from each line,
but we need to maintain them.

Because of the above, we need to trim the final newline, so we don't have a trailing one.
*/
func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	title := readMetaLine(titleSeparator)
	description := readMetaLine(descriptionSeparator)
	tags := strings.Split(readMetaLine(tagSeparator), ", ")

	return Post{
		Title:       title,
		Description: description,
		Tags:        tags,
		Body:        readBody(scanner),
	}, nil
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan() // ignore a line
	buf := bytes.Buffer{}

	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}

	return strings.TrimSuffix(buf.String(), "\n")
}
