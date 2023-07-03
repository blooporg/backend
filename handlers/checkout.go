package handlers

import (
	"fmt"
	"net/http"
)

func CheckoutPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<b>Checkout Page</b>")
}
