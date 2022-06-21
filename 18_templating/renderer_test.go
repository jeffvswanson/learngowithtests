package blogrenderer_test

import (
	"bytes"
	"github.com/jeffvswanson/learngowithtests/18_templating/blogrenderer"
	"testing"
)

func TestRender(t *testing.T) {
	var (
		aPost = blogrenderer.Post{
			Title:       "Hello World",
			Body:        "This is a post",
			Description: "very descriptive",
			Tags:        []string{"go", "tdd"},
		}
	)

	t.Run("convert a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := blogrenderer.Render(&buf, aPost)
		if err != nil {
			t.Fatal(err)
		}
		got := buf.String()
		want := `<h1>Hello World</h1><p>very descriptive</p><ul><li>go</li><li>tdd</li></ul>`
		if got != want {
			t.Errorf("got '%s', want '%s", got, want)
		}
	})
}
