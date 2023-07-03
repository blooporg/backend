package handlers

import (
	"fmt"
	"net/http"
)

func CartPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<b>Cart Page</b>")
}
