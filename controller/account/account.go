package account

import (
	"fmt"
	"net/http"
)

func AccountPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> Account Page!!</h1>")
}
