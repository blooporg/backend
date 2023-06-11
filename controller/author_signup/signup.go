package author_signup

import (
	"net/http"

	"github.com/ibilalkayy/backend/controller/templates"
)

func AuthorSignup(w http.ResponseWriter, r *http.Request) error {
	return templates.AuthorSignupTmpl.Execute(w, nil)
}
