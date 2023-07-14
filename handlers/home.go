package handlers

import (
	"fmt"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	// err := templates.HomeTmpl.Execute(w, nil)
	// if err != nil {
	// 	log.Fatal()
	// }
	fmt.Fprintf(w, "<b>Home Page</b>")
}
