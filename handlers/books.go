package handlers

import (
	"fmt"
	"net/http"
)

func BooksPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<b>Books Page</b>")
}
