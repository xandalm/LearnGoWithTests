package blogrenderer

import (
	"embed"
	"html/template"
	"io"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

var (
	//go:embed "templates/*"
	postTemplate embed.FS
)

type PostRenderer struct {
	templ    *template.Template
	mdParser *parser.Parser
}

func NewRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplate, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)

	return &PostRenderer{templ: templ, mdParser: parser}, nil
}

func (r *PostRenderer) Render(w io.Writer, p Post) error {

	if err := r.templ.ExecuteTemplate(w, "blog.gohtml", newPostVM(p, r)); err != nil {
		return err
	}

	return nil
}

type postViewModel struct {
	Post
	HTMLBody template.HTML
}

func newPostVM(p Post, r *PostRenderer) postViewModel {
	vm := postViewModel{Post: p}
	vm.HTMLBody = template.HTML(markdown.ToHTML([]byte(p.Body), r.mdParser, nil))
	return vm
}
