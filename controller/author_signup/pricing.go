package author_signup

import (
	"net/http"

	"github.com/ibilalkayy/backend/controller/templates"
)

func AuthorPricing(w http.ResponseWriter, r *http.Request) error {
	return templates.AuthorPricingTmpl.Execute(w, nil)
}
