package controller

import (
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) error {
	if r.URL.Path != "/" {
		return PageError.Execute(w, r)
	}
	return IndexTmpl.Execute(w, nil)
}
