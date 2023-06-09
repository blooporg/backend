package account

import (
	"fmt"
	"net/http"
)

func AccountPage(w http.ResponseWriter, r *http.Request) error {
	fmt.Fprintf(w, "<h1> Account Page!!</h1>")
	return nil
}
