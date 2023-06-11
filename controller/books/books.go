package books

import (
	"net/http"

	"github.com/ibilalkayy/backend/controller/templates"
)

func Books(w http.ResponseWriter, r *http.Request) error {
	return templates.BooksTmpl.Execute(w, nil)
}
