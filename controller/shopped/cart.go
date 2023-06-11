package shopped

import (
	"net/http"

	"github.com/ibilalkayy/backend/controller/templates"
)

func Cart(w http.ResponseWriter, r *http.Request) error {
	return templates.CartTmpl.Execute(w, nil)
}
