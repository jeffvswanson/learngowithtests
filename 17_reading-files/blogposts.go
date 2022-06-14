package blogposts

import (
	"bufio"
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

	return Post{Title: title, Description: description, Tags: tags}, nil
}
