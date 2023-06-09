package billing

import (
	"fmt"
	"net/http"
)

func BillingInfo(w http.ResponseWriter, r *http.Request) error {
	fmt.Fprintf(w, "<h1> Billing Info Page!!</h1>")
	return nil
}
