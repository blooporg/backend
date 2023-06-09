package loops

import (
	"fmt"
	"net/http"
)

func Loops(w http.ResponseWriter, r *http.Request) error {
	fmt.Fprintf(w, "<h1> Loops Page!! </h1>")
	return nil
}
