package controller

import (
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) error {
	IndexTmpl.Execute(w, nil)
	return nil
}
