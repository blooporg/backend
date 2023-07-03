package handlers

import (
	"fmt"
	"net/http"
)

func LoopsPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<b>Loops Page</b>")
}
