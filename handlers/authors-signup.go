package handlers

import (
	"fmt"
	"net/http"
)

func AuthorsSignupPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<b>Author Signup Page</b>")
}
