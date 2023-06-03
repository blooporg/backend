package shopped

import (
	"fmt"
	"net/http"
)

func Checkout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> Checkout Page!!</h1>")
}
