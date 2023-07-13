package handlers

import (
	"log"
	"net/http"

	"github.com/ibilalkayy/Bloop/backend/templates"
)

func LoopsPage(w http.ResponseWriter, r *http.Request) {
	err := templates.LoopsTmpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}
