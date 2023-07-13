package handlers

import (
	"log"
	"net/http"

	"github.com/ibilalkayy/Bloop/backend/templates"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	err := templates.HomeTmpl.Execute(w, nil)
	if err != nil {
		log.Fatal()
	}
}
