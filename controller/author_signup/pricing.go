package author_signup

import (
	"fmt"
	"net/http"
)

func AuthorPricing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> Author Pricing Page!!</h1>")
}
