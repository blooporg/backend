package templates

import "html/template"

func MakeTemplate(path string) *template.Template {
	files := []string{path, "view/templates/base.html"}
	return template.Must(template.ParseFiles(files...))
}

var IndexTmpl = MakeTemplate("view/templates/index.html")

var PageError = MakeTemplate("view/templates/page-error.html")

var AccountTmpl = MakeTemplate("view/templates/account/account.html")

var (
	AuthorSignupTmpl  = MakeTemplate("view/templates/author_signup/signup.html")
	AuthorPricingTmpl = MakeTemplate("view/templates/author_signup/pricing.html")
)
