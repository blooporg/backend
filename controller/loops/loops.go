package loops

import (
	"fmt"
	"net/http"
)

func Loops(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> Loops Page!! </h1>")
}
