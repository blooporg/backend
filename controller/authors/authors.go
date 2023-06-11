package authors

import (
	"net/http"

	"github.com/ibilalkayy/backend/controller/templates"
)

func Authors(w http.ResponseWriter, r *http.Request) error {
	return templates.AuthorsTmpl.Execute(w, nil)
}
