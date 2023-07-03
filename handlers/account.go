package handlers

import (
	"fmt"
	"net/http"
)

func AccountPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<b>Account Page</b>")
}
