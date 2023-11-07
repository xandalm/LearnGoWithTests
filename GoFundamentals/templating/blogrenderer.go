package blogrenderer

import (
	"embed"
	"html/template"
	"io"
)

var (
	//go:embed "templates/*"
	postTemplate embed.FS
)

type PostRenderer struct {
	templ *template.Template
}

func NewRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplate, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}
	return &PostRenderer{templ: templ}, nil
}

func (r *PostRenderer) Render(w io.Writer, p Post) error {

	if err := r.templ.ExecuteTemplate(w, "blog.gohtml", p); err != nil {
		return err
	}

	return nil
}
