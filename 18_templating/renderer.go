package blogrenderer

import (
	"bytes"
	"fmt"
	"io"
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

func Render(w io.Writer, p Post) error {
	_, err := fmt.Fprintf(w, "<h1>%s</h1><p>%s</p>%s", p.Title, p.Description, renderTags(p.Tags))
	return err
}

func renderTags(tags []string) string {
	var buf bytes.Buffer
	if len(tags) > 0 {
		buf.WriteString("<ul>")
		for _, tag := range tags {
			buf.WriteString("<li>" + tag + "</li>")
		}
		buf.WriteString("</ul>")
	}
	return buf.String()
}
