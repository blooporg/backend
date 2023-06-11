package controller

import (
	"net/http"

	"github.com/ibilalkayy/backend/controller/templates"
)

func Index(w http.ResponseWriter, r *http.Request) error {
	if r.URL.Path != "/" {
		return templates.PageError.Execute(w, r)
	}
	return templates.IndexTmpl.Execute(w, nil)
}
