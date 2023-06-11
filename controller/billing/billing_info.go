package billing

import (
	"net/http"

	"github.com/ibilalkayy/backend/controller/templates"
)

func BillingInfo(w http.ResponseWriter, r *http.Request) error {
	return templates.BillingInfoTmpl.Execute(w, nil)
}
