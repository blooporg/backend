package handlers

import (
	"fmt"
	"net/http"
)

func LogoutPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<b>Logout Page</b>")
}
