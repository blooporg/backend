package shopped

import (
	"fmt"
	"net/http"
)

func Checkout(w http.ResponseWriter, r *http.Request) error {
	fmt.Fprintf(w, "<h1> Checkout Page!!</h1>")
	return nil
}
