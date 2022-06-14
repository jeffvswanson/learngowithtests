package blogposts_test

import (
	"errors"
	"io/fs"
	"testing"
	"testing/fstest"

	"github.com/jeffvswanson/learngowithtests/17_reading-files/blogposts"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("hi")},
		"hello-world2.md": {Data: []byte("hola")},
	}
	posts, err := blogposts.NewPostsFromFS(fs)
	if err != nil {
		t.Fatal(err)
	}
	if len(posts) != len(fs) {
		t.Errorf("got %d posts, want %d", len(posts), len(fs))
	}
}

type StubFailingFS struct{}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("i purposefully cause your test to fail")
}
