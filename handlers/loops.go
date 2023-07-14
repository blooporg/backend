package handlers

import (
	"fmt"
	"net/http"
)

func LoopsPage(w http.ResponseWriter, r *http.Request) {
	// err := templates.LoopsTmpl.Execute(w, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Fprintf(w, "<b>Loops Page</b>")
}
