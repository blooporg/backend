package handlers

import (
	"fmt"
	"net/http"
)

func PricingPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<b>Pricing Page</b>")
}
