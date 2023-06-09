package authors

import (
	"fmt"
	"net/http"
)

func Authors(w http.ResponseWriter, r *http.Request) error {
	fmt.Fprintf(w, "<h1> Authors Page!!</h1>")
	return nil
}
