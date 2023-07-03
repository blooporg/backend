package handlers

import (
	"fmt"
	"net/http"
)

func SignupPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<b>Signup Page</b>")
}
