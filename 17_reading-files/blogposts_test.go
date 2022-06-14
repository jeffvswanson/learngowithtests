package blogposts_test

import (
	"testing"
	"testing/fstest"

	"github.com/jeffvswanson/learngowithtests/17_reading-files/blogposts"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("hi")},
		"hello-world2.md": {Data: []byte("hola")},
	}
	posts := blogposts.NewPostsFromFS(fs)
	if len(posts) != len(fs) {
		t.Errorf("got %d posts, want %d", len(posts), len(fs))
	}
}
