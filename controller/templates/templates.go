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

var AuthorsTmpl = MakeTemplate("view/templates/authors/authors.html")

var BillingInfoTmpl = MakeTemplate("view/templates/billing/billing_info.html")

var BooksTmpl = MakeTemplate("view/templates/books/books.html")

var LoopsTmpl = MakeTemplate("view/templates/loops/loops.html")

var (
	CartTmpl     = MakeTemplate("view/templates/shopped/cart.html")
	CheckoutTmpl = MakeTemplate("view/templates/shopped/checkout.html")
)
