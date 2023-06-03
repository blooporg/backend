package books

import (
	"fmt"
	"net/http"
)

func Books(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> Books Page!!</h1>")
}
