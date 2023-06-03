package shopped

import (
	"fmt"
	"net/http"
)

func Cart(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> Cart Page!!</h1>")
}
