package account

import (
	"net/http"

	"github.com/ibilalkayy/backend/controller"
)

func AccountPage(w http.ResponseWriter, r *http.Request) error {
	return controller.AccountTmpl.Execute(w, nil)
}
