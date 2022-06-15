package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

const (
	descriptionSeparator = "Description: "
	titleSeparator       = "Title: "
	tagsSeparator        = "Tags: "
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, f := range dir {
		post, err := getPost(fileSystem, f.Name())
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(fileSystem fs.FS, fileName string) (Post, error) {
	postFile, err := fileSystem.Open(fileName)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()
	return newPost(postFile)
}

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readMetaLine := func(metaTag string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), metaTag)
	}

	title := readMetaLine(titleSeparator)                    // strips 'Title: '
	description := readMetaLine(descriptionSeparator)        // strips 'Description: '
	tags := strings.Split(readMetaLine(tagsSeparator), ", ") // strips 'Tags: '

	return Post{Title: title, Description: description, Tags: tags, Body: readBody(scanner)}, nil
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan() // ignore a line, more specifically, the `---` separating the meta information from the body of the post.

	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}
	return strings.TrimSuffix(buf.String(), "\n") // Need to remove the last newline.
}
