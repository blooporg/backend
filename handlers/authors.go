package handlers

import (
	"fmt"
	"net/http"
)

func AuthorsPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<b>Authors Page</b>")
}
