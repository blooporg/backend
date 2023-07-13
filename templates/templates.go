package templates

import "html/template"

// Attaching the base.html file to each template.
func MakeTemplate(path string) *template.Template {
	files := []string{path, "../frontend/html/base.html"}
	return template.Must(template.ParseFiles(files...))
}

// Header Templates
var (
	HomeTmpl  = MakeTemplate("../frontend/html/home.html")
	LoopsTmpl = MakeTemplate("../frontend/html/loops.html")
)
