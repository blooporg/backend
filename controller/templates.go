package controller

import "html/template"

func MakeTemplate(path string) *template.Template {
	files := []string{path, "view/templates/base.html"}
	return template.Must(template.ParseFiles(files...))
}

var (
	IndexTmpl = MakeTemplate("view/templates/index.html")
)
