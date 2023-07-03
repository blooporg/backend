package handlers

import (
	"fmt"
	"net/http"
)

func BillingPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<b>Billing Page</b>")
}
