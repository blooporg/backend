package shopped

import (
	"net/http"

	"github.com/ibilalkayy/backend/controller/templates"
)

func Checkout(w http.ResponseWriter, r *http.Request) error {
	return templates.CheckoutTmpl.Execute(w, nil)
}
